// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package util

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"

	"github.com/elastic/go-ucfg"
	"github.com/elastic/go-ucfg/yaml"
)

type DataSet struct {
	ID             string   `config:"id" json:"id,omitempty" yaml:"id,omitempty"`
	Title          string   `config:"title" json:"title" validate:"required"`
	Release        string   `config:"release" json:"release"`
	Type           string   `config:"type" json:"type" validate:"required"`
	IngestPipeline string   `config:"ingest_pipeline,omitempty" config:"ingest_pipeline" json:"ingest_pipeline,omitempty" yaml:"ingest_pipeline,omitempty"`
	Streams        []Stream `config:"streams" json:"streams,omitempty" yaml:"streams,omitempty" validate:"required"`
	Package        string   `json:"package,omitempty" yaml:"package,omitempty"`

	// Generated fields
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	// Local path to the package dir
	BasePath string `json:"-" yaml:"-"`
}

type Input struct {
	Type        string                   `config:"type" json:"type" validate:"required"`
	Vars        []map[string]interface{} `config:"vars" json:"vars,omitempty" yaml:"vars,omitempty"`
	Title       string                   `config:"title" json:"title,omitempty" yaml:"title,omitempty"`
	Description string                   `config:"description" json:"description,omitempty" yaml:"description,omitempty"`
	Streams     []Stream                 `config:"streams" json:"streams,omitempty" yaml:"streams,omitempty"`
}

type Stream struct {
	Input       string                   `config:"input" json:"input" validate:"required"`
	Vars        []map[string]interface{} `config:"vars" json:"vars,omitempty" yaml:"vars,omitempty"`
	Dataset     string                   `config:"dataset" json:"dataset,omitempty" yaml:"dataset,omitempty"`
	Title       string                   `config:"title" json:"title,omitempty" yaml:"title,omitempty"`
	Description string                   `config:"description" json:"description,omitempty" yaml:"description,omitempty"`
}

func NewDataset(basePath string, p *Package) (*DataSet, error) {

	// Check if manifest exists
	manifestPath := filepath.Join(basePath, "manifest.yml")
	_, err := os.Stat(manifestPath)
	if err != nil && os.IsNotExist(err) {
		return nil, errors.Wrapf(err, "manifest does not exist for package: %s", p.BasePath)
	}

	datasetPath := filepath.Base(basePath)

	manifest, err := yaml.NewConfigWithFile(manifestPath, ucfg.PathSep("."))
	var d = &DataSet{
		Package: p.Name,
		// This is the name of the directory of the dataset
		Path:     datasetPath,
		BasePath: basePath,
	}
	// go-ucfg automatically calls the `Validate` method on the Dataset object here
	err = manifest.Unpack(d)
	if err != nil {
		return nil, errors.Wrapf(err, "error building dataset (path: %s) in package: %s", datasetPath, p.Name)
	}

	// if id is not set, {package}.{datasetPath} is the default
	if d.ID == "" {
		d.ID = p.Name + "." + datasetPath
	}

	if d.Release == "" {
		d.Release = "beta"
	}

	return d, nil
}

func (d *DataSet) Validate() error {
	pipelineDir := filepath.Join(d.BasePath, "elasticsearch", "ingest-pipeline")
	paths, err := filepath.Glob(filepath.Join(pipelineDir, "*"))
	if err != nil {
		return err
	}

	if strings.Contains(d.ID, "-") {
		return fmt.Errorf("dataset name is not allowed to contain `-`: %s", d.ID)
	}

	if d.IngestPipeline == "" {
		// Check that no ingest pipeline exists in the directory except default
		for _, path := range paths {
			if filepath.Base(path) == "default.json" || filepath.Base(path) == "default.yml" {
				d.IngestPipeline = "default"
				break
			}
		}
	}

	if d.IngestPipeline == "" && len(paths) > 0 {
		return fmt.Errorf("Package contains pipelines which are not used: %v, %s", paths, d.ID)
	}

	// In case an ingest pipeline is set, check if it is around
	if d.IngestPipeline != "" {
		_, errJSON := os.Stat(filepath.Join(pipelineDir, d.IngestPipeline+".json"))
		_, errYAML := os.Stat(filepath.Join(pipelineDir, d.IngestPipeline+".yml"))

		if os.IsNotExist(errYAML) && os.IsNotExist(errJSON) {
			return fmt.Errorf("Defined ingest_pipeline does not exist: %s", pipelineDir+d.IngestPipeline)
		}
	}
	return nil
}
