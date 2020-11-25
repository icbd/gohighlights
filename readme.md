# GoHighlights

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/icbd/gohighlights/CI)
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/icbd/gohighlights)

This is GoHighlights's server project, the frontend project is [https://github.com/icbd/gohighlights_ext](https://github.com/icbd/gohighlights_ext) .

## Features

* Add Highlight
* Change Color
* Remove Highlight
* Replay Highlight

![https://github.com/icbd/gohighlights_ext/blob/master/demo.png](https://github.com/icbd/gohighlights_ext/blob/master/demo.png)

## Config

### ENV

ENV Tag|Description|Default
---|---|---
GIN_MODE | gin mode (debug/test/release) | `debug` 
DB_TYPE | database type (mysql/sqlite) | `sqlite`
DB_DSN | database data source name | `root:password@tcp(127.0.0.1:3306)/dbname`
CONF_LOC | config file location | `./config.yaml`
ES_URL | elasticsearch url | `http://localhost:9200`

### CMD

```shell script
GIN_MODE=debug DB_TYPE="" DB_DSN="" cmd -db=migrate
```

## Run test

including sub-packages

```shell script
GIN_MODE=test go run ./bin/migrate_cmd.go -db=migrate
GIN_MODE=test go test ./... -v
```

## License

MIT, see [LICENSE](LICENSE)