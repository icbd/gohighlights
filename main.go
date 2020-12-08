package main

import (
	"github.com/gin-gonic/autotls"
	"github.com/icbd/gohighlights/config"
	"github.com/icbd/gohighlights/indices"
	"github.com/icbd/gohighlights/models"
	"github.com/icbd/gohighlights/models/migrations"
	"github.com/icbd/gohighlights/routes"
	mgr "github.com/icbd/gorm-migration"
	"github.com/spf13/cast"
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
	indices.Ping()

	domains := cast.ToStringSlice(config.Get("app.domains"))
	if domains == nil {
		log.Fatal(server.ListenAndServe())
	} else {
		log.Fatal(autotls.Run(server.Handler, domains...))
	}
}
