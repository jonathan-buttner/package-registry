# MySQL Integration

This integration periodically fetches logs and metrics from [https://www.mysql.com/](MySQL) servers.

## Compatibility

The `error` and `slowlog` datasets were tested with logs from MySQL 5.5, 5.7 and 8.0, MariaDB 10.1, 10.2 and 10.3, and Percona 5.7 and 8.0.

The `galera_status` and `status` datasets were tested with MySQL and Percona 5.7 and 8.0 and are expected to work with all
versions >= 5.7.0. It is also tested with MariaDB 10.2, 10.3 and 10.4.

## Logs

### error

The `error` dataset collects the MySQL error logs.

**Exported fields**

| Field | Description | Type |
|---|---|---|
| mysql.thread_id | The connection or thread ID for the query. | long |


### slowlog

The `slowlog` dataset collects the MySQL slow logs.

**Exported fields**

| Field | Description | Type |
|---|---|---|
| mysql.slowlog.bytes_received | The number of bytes received from client. | long |
| mysql.slowlog.bytes_sent | The number of bytes sent to client. | long |
| mysql.slowlog.current_user | Current authenticated user, used to determine access privileges. Can differ from the value for user. | keyword |
| mysql.slowlog.filesort | Whether filesort optimization was used. | boolean |
| mysql.slowlog.filesort_on_disk | Whether filesort optimization was used and it needed temporary tables on disk. | boolean |
| mysql.slowlog.full_join | Whether a full join was needed for the slow query (no indexes were used for joins). | boolean |
| mysql.slowlog.full_scan | Whether a full table scan was needed for the slow query. | boolean |
| mysql.slowlog.innodb.io_r_bytes | Bytes read during page read operations. | long |
| mysql.slowlog.innodb.io_r_ops | Number of page read operations. | long |
| mysql.slowlog.innodb.io_r_wait.sec | How long it took to read all needed data from storage. | long |
| mysql.slowlog.innodb.pages_distinct | Approximated count of pages accessed to execute the query. | long |
| mysql.slowlog.innodb.queue_wait.sec | How long the query waited to enter the InnoDB queue and to be executed once in the queue. | long |
| mysql.slowlog.innodb.rec_lock_wait.sec | How long the query waited for locks. | long |
| mysql.slowlog.innodb.trx_id | Transaction ID | keyword |
| mysql.slowlog.killed | Code of the reason if the query was killed. | keyword |
| mysql.slowlog.last_errno | Last SQL error seen. | keyword |
| mysql.slowlog.lock_time.sec | The amount of time the query waited for the lock to be available. The value is in seconds, as a floating point number. | float |
| mysql.slowlog.log_slow_rate_limit | Slow log rate limit, a value of 100 means that one in a hundred queries or sessions are being logged. | keyword |
| mysql.slowlog.log_slow_rate_type | Type of slow log rate limit, it can be `session` if the rate limit is applied per session, or `query` if it applies per query. | keyword |
| mysql.slowlog.merge_passes | Number of merge passes executed for the query. | long |
| mysql.slowlog.priority_queue | Whether a priority queue was used for filesort. | boolean |
| mysql.slowlog.query | The slow query. | keyword |
| mysql.slowlog.query_cache_hit | Whether the query cache was hit. | boolean |
| mysql.slowlog.read_first | The number of times the first entry in an index was read. | long |
| mysql.slowlog.read_key | The number of requests to read a row based on a key. | long |
| mysql.slowlog.read_last | The number of times the last key in an index was read. | long |
| mysql.slowlog.read_next | The number of requests to read the next row in key order. | long |
| mysql.slowlog.read_prev | The number of requests to read the previous row in key order. | long |
| mysql.slowlog.read_rnd | The number of requests to read a row based on a fixed position. | long |
| mysql.slowlog.read_rnd_next | The number of requests to read the next row in the data file. | long |
| mysql.slowlog.rows_affected | The number of rows modified by the query. | long |
| mysql.slowlog.rows_examined | The number of rows scanned by the query. | long |
| mysql.slowlog.rows_sent | The number of rows returned by the query. | long |
| mysql.slowlog.schema | The schema where the slow query was executed. | keyword |
| mysql.slowlog.sort_merge_passes | Number of merge passes that the sort algorithm has had to do. | long |
| mysql.slowlog.sort_range_count | Number of sorts that were done using ranges. | long |
| mysql.slowlog.sort_rows | Number of sorted rows. | long |
| mysql.slowlog.sort_scan_count | Number of sorts that were done by scanning the table. | long |
| mysql.slowlog.tmp_disk_tables | Number of temporary tables created on disk for this query. | long |
| mysql.slowlog.tmp_table | Whether a temporary table was used to resolve the query. | boolean |
| mysql.slowlog.tmp_table_on_disk | Whether the query needed temporary tables on disk. | boolean |
| mysql.slowlog.tmp_table_sizes | Size of temporary tables created for this query. | long |
| mysql.slowlog.tmp_tables | Number of temporary tables created for this query | long |
| mysql.thread_id | The connection or thread ID for the query. | long |


## Metrics

### galera_status

The `galera_status` dataset periodically fetches metrics from [http://galeracluster.com/](Galera)-MySQL cluster servers.

An example event for `galera_status` looks as following:

```$json
{
    "@timestamp":"2016-05-23T08:05:34.853Z",
    "agent": {
        "hostname": "host.example.com",
        "name": "host.example.com"
    },
    "event": {
        "dataset": "mysql.galera_status",
        "duration": 115000
    },
    "metricset": {
        "name": "galera_status"
    },
    "mysql":{
        "galera_status":{
            "apply": {
                "oooe": 0,
                "oool": 0,
                "window": 1
            },
            "connected": "ON",
            "flow_ctl": {
                "recv": 0,
                "sent": 0,
                "paused": 0,
                "paused_ns": 0
            },
            "ready": "ON",
            "received": {
                "count": 173,
                "bytes": 152425
            },
            "local": {
                "state": "Synced",
                "bf_aborts": 0,
                "cert_failures": 0,
                "commits": 1325,
                "recv": {
                    "queue_max": 2,
                    "queue_min": 0,
                    "queue": 0,
                    "queue_avg": 0.011561
                },
                "replays": 0,
                "send": {
                    "queue_min": 0,
                    "queue": 0,
                    "queue_avg": 0,
                    "queue_max": 1
                }
            },
            "evs": {
                "evict": "",
                "state": "OPERATIONAL"
            },
            "repl": {
                "bytes": 1689804,
                "data_bytes": 1540647,
                "keys": 4170,
                "keys_bytes": 63973,
                "other_bytes": 0,
                "count": 1331
            },
            "commit": {
                "oooe": 0,
                "window": 1
            },
            "cluster": {
                "conf_id": 930,
                "size": 3,
                "status": "Primary"
            },
            "last_committed": 23944,
            "cert": {
                "deps_distance": 43.524557,
                "index_size": 22,
                "interval": 0
            }
        }
    }
}
```

The fields reported are:

**Exported fields**

| Field | Description | Type |
|---|---|---|
| mysql.galera_status.apply.oooe | How often applier started write-set applying out-of-order (parallelization efficiency). | double |
| mysql.galera_status.apply.oool | How often write-set was so slow to apply that write-set with higher seqno's were applied earlier. Values closer to 0 refer to a greater gap between slow and fast write-sets. | double |
| mysql.galera_status.apply.window | Average distance between highest and lowest concurrently applied seqno. | double |
| mysql.galera_status.cert.deps_distance | Average distance between highest and lowest seqno value that can be possibly applied in parallel (potential degree of parallelization). | double |
| mysql.galera_status.cert.index_size | The number of entries in the certification index. | long |
| mysql.galera_status.cert.interval | Average number of transactions received while a transaction replicates. | double |
| mysql.galera_status.cluster.conf_id | Total number of cluster membership changes happened. | long |
| mysql.galera_status.cluster.size | Current number of members in the cluster. | long |
| mysql.galera_status.cluster.status | Status of this cluster component. That is, whether the node is part of a PRIMARY or NON_PRIMARY component. | keyword |
| mysql.galera_status.commit.oooe | How often a transaction was committed out of order. | double |
| mysql.galera_status.commit.window | Average distance between highest and lowest concurrently committed seqno. | long |
| mysql.galera_status.connected | If the value is OFF, the node has not yet connected to any of the cluster components. This may be due to misconfiguration. Check the error log for proper diagnostics. | keyword |
| mysql.galera_status.evs.evict | Lists the UUID's of all nodes evicted from the cluster. Evicted nodes cannot rejoin the cluster until you restart their mysqld processes. | keyword |
| mysql.galera_status.evs.state | Shows the internal state of the EVS Protocol. | keyword |
| mysql.galera_status.flow_ctl.paused | The fraction of time since the last FLUSH STATUS command that replication was paused due to flow control. In other words, how much the slave lag is slowing down the cluster. | double |
| mysql.galera_status.flow_ctl.paused_ns | The total time spent in a paused state measured in nanoseconds. | long |
| mysql.galera_status.flow_ctl.recv | Returns the number of FC_PAUSE events the node has received, including those the node has sent. Unlike most status variables, the counter for this one does not reset every time you run the query. | long |
| mysql.galera_status.flow_ctl.sent | Returns the number of FC_PAUSE events the node has sent. Unlike most status variables, the counter for this one does not reset every time you run the query. | long |
| mysql.galera_status.last_committed | The sequence number, or seqno, of the last committed transaction. | long |
| mysql.galera_status.local.bf_aborts | Total number of local transactions that were aborted by slave transactions while in execution. | long |
| mysql.galera_status.local.cert_failures | Total number of local transactions that failed certification test. | long |
| mysql.galera_status.local.commits | Total number of local transactions committed. | long |
| mysql.galera_status.local.recv.queue | Current (instantaneous) length of the recv queue. | long |
| mysql.galera_status.local.recv.queue_avg | Recv queue length averaged over interval since the last FLUSH STATUS command. Values considerably larger than 0.0 mean that the node cannot apply write-sets as fast as they are received and will generate a lot of replication throttling. | double |
| mysql.galera_status.local.recv.queue_max | The maximum length of the recv queue since the last FLUSH STATUS command. | long |
| mysql.galera_status.local.recv.queue_min | The minimum length of the recv queue since the last FLUSH STATUS command. | long |
| mysql.galera_status.local.replays | Total number of transaction replays due to asymmetric lock granularity. | long |
| mysql.galera_status.local.send.queue | Current (instantaneous) length of the send queue. | long |
| mysql.galera_status.local.send.queue_avg | Send queue length averaged over time since the last FLUSH STATUS command. Values considerably larger than 0.0 indicate replication throttling or network throughput issue. | double |
| mysql.galera_status.local.send.queue_max | The maximum length of the send queue since the last FLUSH STATUS command. | long |
| mysql.galera_status.local.send.queue_min | The minimum length of the send queue since the last FLUSH STATUS command. | long |
| mysql.galera_status.local.state | Internal Galera Cluster FSM state number. | keyword |
| mysql.galera_status.ready | Whether the server is ready to accept queries. | keyword |
| mysql.galera_status.received.bytes | Total size of write-sets received from other nodes. | long |
| mysql.galera_status.received.count | Total number of write-sets received from other nodes. | long |
| mysql.galera_status.repl.bytes | Total size of write-sets replicated. | long |
| mysql.galera_status.repl.count | Total number of write-sets replicated (sent to other nodes). | long |
| mysql.galera_status.repl.data_bytes | Total size of data replicated. | long |
| mysql.galera_status.repl.keys | Total number of keys replicated. | long |
| mysql.galera_status.repl.keys_bytes | Total size of keys replicated. | long |
| mysql.galera_status.repl.other_bytes | Total size of other bits replicated. | long |


### status

The MySQL `status` dataset collects data from MySQL by running a `SHOW GLOBAL STATUS;` SQL query. This query returns a large number of metrics.

An example event for `status` looks as following:

```$json
{
    "@timestamp":"2016-05-23T08:05:34.853Z",
    "agent": {
        "hostname": "host.example.com",
        "name": "host.example.com"
    },
    "event": {
        "dataset": "mysql.status",
        "duration": 115000
    },
    "metricset": {
        "name": "status"
    },
    "mysql": {
        "status": {
            "aborted": {
                "clients": 3,
                "connects": 4
            },
            "binlog": {
                "cache": {
                    "disk_use": 0,
                    "use": 0
                }
            },
            "bytes": {
                "received": 1272,
                "sent": 47735
            },
            "command": {
                "delete": 0,
                "insert": 0,
                "select": 1,
                "update": 0
            },
            "connections": 12,
            "created": {
                "tmp": {
                    "disk_tables": 0,
                    "files": 5,
                    "tables": 6
                }
            },
            "delayed": {
                "errors": 0,
                "insert_threads": 0,
                "writes": 0
            },
            "flush_commands": 1,
            "handler": {
                "commit": 0,
                "delete": 0,
                "external_lock": 140,
                "mrr_init": 0,
                "prepare": 0,
                "read": {
                    "first": 3,
                    "key": 2,
                    "last": 0,
                    "next": 32,
                    "prev": 0,
                    "rnd": 0,
                    "rnd_next": 1728
                },
                "rollback": 0,
                "savepoint": 0,
                "savepoint_rollback": 0,
                "update": 0,
                "write": 1705
            },
            "innodb": {
                "buffer_pool": {
                    "bytes": {
                        "data": 6914048,
                        "dirty": 0
                    },
                    "pages": {
                        "data": 422,
                        "dirty": 0,
                        "flushed": 207,
                        "free": 7768,
                        "misc": 1,
                        "total": 8191
                    },
                    "pool": {
                        "reads": 423,
                        "wait_free": 0
                    },
                    "read": {
                        "ahead": 0,
                        "ahead_evicted": 0,
                        "ahead_rnd": 0,
                        "requests": 14198
                    },
                    "write_requests": 207
                }
            },
            "max_used_connections": 3,
            "open": {
                "files": 16,
                "streams": 0,
                "tables": 60
            },
            "opened_tables": 67,
            "queries": 10,
            "questions": 9,
            "threads": {
                "cached": 0,
                "connected": 3,
                "created": 3,
                "running": 1
            }
        }
    }
}
```

The fields reported are:

**Exported fields**

| Field | Description | Type |
|---|---|---|
| mysql.status.aborted.clients | The number of connections that were aborted because the client died without closing the connection properly. | long |
| mysql.status.aborted.connects | The number of failed attempts to connect to the MySQL server. | long |
| mysql.status.binlog.cache.disk_use |  | long |
| mysql.status.binlog.cache.use |  | long |
| mysql.status.bytes.received | The number of bytes received from all clients. | long |
| mysql.status.bytes.sent | The number of bytes sent to all clients. | long |
| mysql.status.command.delete | The number of DELETE queries since startup. | long |
| mysql.status.command.insert | The number of INSERT queries since startup. | long |
| mysql.status.command.select | The number of SELECT queries since startup. | long |
| mysql.status.command.update | The number of UPDATE queries since startup. | long |
| mysql.status.connections |  | long |
| mysql.status.created.tmp.disk_tables |  | long |
| mysql.status.created.tmp.files |  | long |
| mysql.status.created.tmp.tables |  | long |
| mysql.status.delayed.errors |  | long |
| mysql.status.delayed.insert_threads |  | long |
| mysql.status.delayed.writes |  | long |
| mysql.status.flush_commands |  | long |
| mysql.status.handler.commit | The number of internal COMMIT statements. | long |
| mysql.status.handler.delete | The number of times that rows have been deleted from tables. | long |
| mysql.status.handler.external_lock | The server increments this variable for each call to its external_lock() function, which generally occurs at the beginning and end of access to a table instance. | long |
| mysql.status.handler.mrr_init | The number of times the server uses a storage engine's own Multi-Range Read implementation for table access. | long |
| mysql.status.handler.prepare | A counter for the prepare phase of two-phase commit operations. | long |
| mysql.status.handler.read.first | The number of times the first entry in an index was read. | long |
| mysql.status.handler.read.key | The number of requests to read a row based on a key. | long |
| mysql.status.handler.read.last | The number of requests to read the last key in an index. | long |
| mysql.status.handler.read.next | The number of requests to read the next row in key order. | long |
| mysql.status.handler.read.prev | The number of requests to read the previous row in key order. | long |
| mysql.status.handler.read.rnd | The number of requests to read a row based on a fixed position. | long |
| mysql.status.handler.read.rnd_next | The number of requests to read the next row in the data file. | long |
| mysql.status.handler.rollback | The number of requests for a storage engine to perform a rollback operation. | long |
| mysql.status.handler.savepoint | The number of requests for a storage engine to place a savepoint. | long |
| mysql.status.handler.savepoint_rollback | The number of requests for a storage engine to roll back to a savepoint. | long |
| mysql.status.handler.update | The number of requests to update a row in a table. | long |
| mysql.status.handler.write | The number of requests to insert a row in a table. | long |
| mysql.status.innodb.buffer_pool.bytes.data | The total number of bytes in the InnoDB buffer pool containing data. | long |
| mysql.status.innodb.buffer_pool.bytes.dirty | The total current number of bytes held in dirty pages in the InnoDB buffer pool. | long |
| mysql.status.innodb.buffer_pool.dump_status | The progress of an operation to record the pages held in the InnoDB buffer pool, triggered by the setting of innodb_buffer_pool_dump_at_shutdown or innodb_buffer_pool_dump_now. | long |
| mysql.status.innodb.buffer_pool.load_status | The progress of an operation to warm up the InnoDB buffer pool by reading in a set of pages corresponding to an earlier point in time, triggered by the setting of innodb_buffer_pool_load_at_startup or innodb_buffer_pool_load_now. | long |
| mysql.status.innodb.buffer_pool.pages.data | he number of pages in the InnoDB buffer pool containing data. | long |
| mysql.status.innodb.buffer_pool.pages.dirty | The current number of dirty pages in the InnoDB buffer pool. | long |
| mysql.status.innodb.buffer_pool.pages.flushed | The number of requests to flush pages from the InnoDB buffer pool. | long |
| mysql.status.innodb.buffer_pool.pages.free | The number of free pages in the InnoDB buffer pool. | long |
| mysql.status.innodb.buffer_pool.pages.latched | The number of latched pages in the InnoDB buffer pool. | long |
| mysql.status.innodb.buffer_pool.pages.misc | The number of pages in the InnoDB buffer pool that are busy because they have been allocated for administrative overhead, such as row locks or the adaptive hash index. | long |
| mysql.status.innodb.buffer_pool.pages.total | The total size of the InnoDB buffer pool, in pages. | long |
| mysql.status.innodb.buffer_pool.pool.reads | The number of logical reads that InnoDB could not satisfy from the buffer pool, and had to read directly from disk. | long |
| mysql.status.innodb.buffer_pool.pool.resize_status | The status of an operation to resize the InnoDB buffer pool dynamically, triggered by setting the innodb_buffer_pool_size parameter dynamically. | long |
| mysql.status.innodb.buffer_pool.pool.wait_free | Normally, writes to the InnoDB buffer pool happen in the background. When InnoDB needs to read or create a page and no clean pages are available, InnoDB flushes some dirty pages first and waits for that operation to finish. This counter counts instances of these waits. | long |
| mysql.status.innodb.buffer_pool.read.ahead | The number of pages read into the InnoDB buffer pool by the read-ahead background thread. | long |
| mysql.status.innodb.buffer_pool.read.ahead_evicted | The number of pages read into the InnoDB buffer pool by the read-ahead background thread that were subsequently evicted without having been accessed by queries. | long |
| mysql.status.innodb.buffer_pool.read.ahead_rnd | The number of "random" read-aheads initiated by InnoDB. | long |
| mysql.status.innodb.buffer_pool.read.requests | The number of logical read requests. | long |
| mysql.status.innodb.buffer_pool.write_requests | The number of writes done to the InnoDB buffer pool. | long |
| mysql.status.max_used_connections |  | long |
| mysql.status.open.files |  | long |
| mysql.status.open.streams |  | long |
| mysql.status.open.tables |  | long |
| mysql.status.opened_tables |  | long |
| mysql.status.queries | The number of statements executed by the server. This variable includes statements executed within stored programs, unlike the Questions variable. It does not count COM_PING or COM_STATISTICS commands. | long |
| mysql.status.questions | The number of statements executed by the server. This includes only statements sent to the server by clients and not statements executed within stored programs, unlike the Queries variable. This variable does not count COM_PING, COM_STATISTICS, COM_STMT_PREPARE, COM_STMT_CLOSE, or COM_STMT_RESET commands. | long |
| mysql.status.threads.cached | The number of cached threads. | long |
| mysql.status.threads.connected | The number of connected threads. | long |
| mysql.status.threads.created | The number of created threads. | long |
| mysql.status.threads.running | The number of running threads. | long |

