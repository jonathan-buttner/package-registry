// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/blang/semver"

	"github.com/elastic/package-registry/util"
)

func searchHandler(packagesBasePath string, cacheTime time.Duration) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		var kibanaVersion *semver.Version
		var category string
		// Leaving out `a` here to not use a reserved name
		var packageQuery string
		var all bool
		var internal bool
		var err error

		// Read query filter params which can affect the output
		if len(query) > 0 {
			if v := query.Get("kibana"); v != "" {
				kibanaVersion, err = semver.New(v)
				if err != nil {
					notFound(w, fmt.Errorf("invalid Kibana version '%s': %s", v, err))
					return
				}
			}

			if v := query.Get("category"); v != "" {
				if v != "" {
					category = v
				}
			}

			if v := query.Get("package"); v != "" {
				if v != "" {
					packageQuery = v
				}
			}

			if v := query.Get("all"); v != "" {
				if v != "" {
					// Default is false, also on error
					all, _ = strconv.ParseBool(v)
				}
			}

			if v := query.Get("internal"); v != "" {
				if v != "" {
					// In case of error, keep it false
					internal, _ = strconv.ParseBool(v)
				}
			}
		}

		packages, err := util.GetPackages(packagesBasePath)
		if err != nil {
			notFound(w, fmt.Errorf("problem fetching packages: %s", err))
			return
		}
		packagesList := map[string]map[string]util.Package{}

		// Checks that only the most recent version of an integration is added to the list
		for _, p := range packages {

			// Skip internal packages by default
			if p.Internal && !internal {
				continue
			}

			// Filter by category first as this could heavily reduce the number of packages
			// It must happen before the version filtering as there only the newest version
			// is exposed and there could be an older package with more versions.
			if category != "" && !p.HasCategory(category) {
				continue
			}

			if kibanaVersion != nil {
				if valid := p.HasKibanaVersion(kibanaVersion); !valid {
					continue
				}
			}

			// If package Query is set, all versions of this package are returned
			if packageQuery != "" && packageQuery != p.Name {
				continue
			}

			if !all {
				// Check if the version exists and if it should be added or not.
				for _, versions := range packagesList {
					for _, pp := range versions {
						if pp.Name == p.Name {

							// If the package in the list is newer, do nothing. Otherwise delete and later add the new one.
							if pp.IsNewer(p) {
								continue
							}

							delete(packagesList[pp.Name], pp.Version)
						}
					}
				}
			}

			if _, ok := packagesList[p.Name]; !ok {
				packagesList[p.Name] = map[string]util.Package{}
			}
			packagesList[p.Name][p.Version] = p
		}

		data, err := getPackageOutput(packagesList)
		if err != nil {
			notFound(w, err)
			return
		}

		cacheHeaders(w, cacheTime)
		jsonHeader(w)
		fmt.Fprint(w, string(data))
	}
}

func getPackageOutput(packagesList map[string]map[string]util.Package) ([]byte, error) {

	separator := "@"
	// Packages need to be sorted to be always outputted in the same order
	var keys []string
	for key, k := range packagesList {
		for v := range k {
			keys = append(keys, key+separator+v)
		}
	}
	sort.Strings(keys)

	var output []map[string]interface{}

	for _, k := range keys {
		parts := strings.Split(k, separator)
		m := packagesList[parts[0]][parts[1]]
		data := map[string]interface{}{
			"name":        m.Name,
			"description": m.Description,
			"version":     m.Version,
			"type":        m.Type,
			"download":    m.GetDownloadPath(),
			"path":        m.GetUrlPath(),
		}
		if m.Title != nil {
			data["title"] = *m.Title
		}
		if m.Icons != nil {
			data["icons"] = m.Icons
		}
		if m.Internal {
			data["internal"] = true
		}
		output = append(output, data)
	}

	// Instead of return `null` in case of an empty array, return []
	if len(output) == 0 {
		return []byte("[]"), nil
	}

	return json.MarshalIndent(output, "", "  ")
}
