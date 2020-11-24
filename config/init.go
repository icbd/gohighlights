package config

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/icbd/gohighlights/utils"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

const (
	EnvKeyConfigLocation = "CONF_LOC"
	DefaultConfigFile    = "./config.yaml"
)

func init() {
	parseConfigFile(configFileLocation())
}

func configFileLocation() string {
	loc, ok := os.LookupEnv(EnvKeyConfigLocation)
	if !ok {
		loc = DefaultConfigFile
	}

	if !filepath.IsAbs(loc) {
		loc = filepath.Join(utils.RootPath(), loc)
	}
	return loc
}

func parseConfigFile(file string) {
	var configBuffer bytes.Buffer
	//tmpl := template.Must(template.New(filepath.Base(file)).Funcs(template.FuncMap{"ENV": func(k string) string { return os.Getenv(k) }}).ParseFiles(file))
	tmpl := template.New(filepath.Base(file))
	tmpl.Funcs(template.FuncMap{"ENV": tmplENV})
	template.Must(tmpl.ParseFiles(file))
	if err := tmpl.Execute(&configBuffer, nil); err != nil {
		log.Fatal(err)
	}
	viper.SetConfigType("yaml")
	if err := viper.ReadConfig(&configBuffer); err != nil {
		log.Fatal(err)
	}
}

// tmplENV template function
// EMV "v1" # return os.Getenv("v1")
// EMV "v1" "v2" # if os.Getenv("v1") blank, return "v2"
// ENV # invalid params return ""
func tmplENV(keys ...string) (value string) {
	switch len(keys) {
	case 1:
		value = os.Getenv(keys[0])
	case 2:
		value = os.Getenv(keys[0])
		if value == "" {
			value = keys[1]
		}
	}
	return
}

func Get(key string) interface{} {
	return viper.Get(configKey(key))
}

func GetString(key string) string {
	return viper.GetString(configKey(key))
}

func GetInt(key string) int {
	return viper.GetInt(configKey(key))
}

func GetBool(key string) bool {
	return viper.GetBool(configKey(key))
}

func Set(key string, value interface{}) {
	viper.Set(configKey(key), value)
}

// configKey trans key to local key
func configKey(key string) string {
	return gin.Mode() + "." + key
}
