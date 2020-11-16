package main

import (
	"flag"
	"github.com/icbd/gohighlights/models"
	"github.com/icbd/gohighlights/models/migrations"
	mgr "github.com/icbd/gorm-migration"
	"log"
)

var migrateType mgr.MigrateType

func main() {
	flag.Var(&migrateType, "db", "-db=migrate")
	flag.Parse()

	if err := models.Ping(); err != nil {
		log.Fatal(err)
	}

	migrations.Migrate(migrateType)
}
