# GoHighlights

## Config

### ENV

```text
GIN_MODE: gin mode. default: debug. debug / test/ release
DB_TYPE: database type. mysql / sqlite
DB_DSN: database data source name. `root:password@tcp(127.0.0.1:3306)/dbname`
CONF_LOC: config file location. `./config.yaml`
```

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