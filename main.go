package main

import (
	"github.com/icbd/gohighlights/config"
	"github.com/icbd/gohighlights/models"
	"github.com/icbd/gohighlights/models/migrations"
	"github.com/icbd/gohighlights/routes"
	mgr "github.com/icbd/gorm-migration"
	"log"
	"net/http"
)

func main() {
	server := &http.Server{
		Handler: routes.R(),
		Addr:    config.GetString("app.addr"),
	}

	if err := models.Ping(); err != nil {
		log.Fatal(err)
	}

	migrations.Migrate(mgr.Check)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
