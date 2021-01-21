# GoHighlights

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/icbd/gohighlights/CI)
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/icbd/gohighlights)

This is GoHighlights's server project, the frontend project is [https://github.com/icbd/gohighlights_ext](https://github.com/icbd/gohighlights_ext) .

## Features

### Ready

- [x] Add Highlight
- [x] Change Color
- [x] Remove Highlight
- [x] Replay Highlight

### Todo

- [ ] History Dashboard
- [ ] User-friendly Personalization Setting
- [ ] Web Timer
- [ ] Reading Report
- [ ] Read Later Box

![https://github.com/icbd/gohighlights_ext/blob/master/demo.png](https://github.com/icbd/gohighlights_ext/blob/master/demo.png)

You can try the online version of the chrome store: 

> [https://chrome.google.com/webstore/detail/go-highlights/homlcfpinafhealhlmjkmdjdejppmmlk](https://chrome.google.com/webstore/detail/go-highlights/homlcfpinafhealhlmjkmdjdejppmmlk)

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

0. Edit Config File

```shell script
vi config.yaml
```

If you are using MySQL, please create the  database manually.

1. Migration

```shell script
GIN_MODE=debug go run ./bin/migrate_cmd.go -db=migrate
```

Also see [https://github.com/icbd/gorm-migration](https://github.com/icbd/gorm-migration) .

2. Run Server

```shell script
GIN_MODE=debug go run ./main.go
```

## Run test

including sub-packages

```shell script
GIN_MODE=test go run ./bin/migrate_cmd.go -db=migrate
GIN_MODE=test go test ./... -v
```

## License

MIT, see [LICENSE](LICENSE)