package models

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/icbd/gohighlights/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var db *gorm.DB
var _validator = binding.Validator.Engine().(*validator.Validate)

func init() {
	var err error
	dbDsn := config.GetString("db.dsn")
	dbType := config.GetString("db.type")
	if gin.Mode() != gin.ReleaseMode {
		log.Printf("dbType:[%s]\tdbDsn:[%s]\n", dbType, dbDsn)
	}

	switch dbType {
	case "mysql":
		db, err = gorm.Open(mysql.Open(dbDsn), dbConfig())
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(dbDsn), dbConfig())
	default:
		log.Fatalf("Database not support yet: %s", dbType)
	}
	if err != nil {
		log.Fatalf("Database open failed: %s", err)
	}

	if err := Ping(); err != nil {
		log.Fatalf("Database ping test failed: %s", err)
	}

	if err := RegisterGormCallBack(); err != nil {
		log.Fatalf("RegisterGormCallBack: %s", err)
	}
}

func dbConfig() *gorm.Config {
	c := gorm.Config{}
	if gin.Mode() != gin.ReleaseMode {
		c.Logger = logger.Default.LogMode(logger.Info)
	}
	return &c
}
