package indices

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/icbd/gohighlights/config"
	"github.com/olivere/elastic/v7"
	"log"
)

var client *elastic.Client
var Enable bool // if not, jump ES processing
var indexNamePrefix string
var NotEnabledError = fmt.Errorf("es is not enabled yet")

func init() {
	indexNamePrefix = config.GetString("es.prefix")
	Enable = config.GetBool("es.enable")
	if !Enable {
		return
	}

	var err error
	esConfig := []elastic.ClientOptionFunc{elastic.SetURL(config.GetString("es.url"))}
	if gin.Mode() != gin.ReleaseMode {
		esConfig = append(esConfig, elastic.SetTraceLog(log.New(log.Writer(), "\n", 0)))
	}
	client, err = elastic.NewClient(esConfig...)
	if err != nil {
		log.Fatal(err)
	}
}

func Client() *elastic.Client {
	return client
}

func IndexName(name string) string {
	return indexNamePrefix + name
}

type SetupMethod func() error

func setupMethods() []SetupMethod {
	return []SetupMethod{
		SetupMarkIndex,
	}
}

// Ping and init indices
func Ping() {
	if !Enable {
		return
	}

	info, _, err := Client().Ping(config.GetString("es.url")).Do(context.Background())

	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Elasticsearch Ping: %#v\n", info)
	}

	for _, f := range setupMethods() {
		if err := f(); err != nil {
			log.Fatal(err)
		}
	}
}
