package indices

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/icbd/gohighlights/config"
	"github.com/olivere/elastic/v7"
	"log"
)

var Client *elastic.Client
var Use bool // if not, jump ES processing

func init() {
	Use = config.GetBool("es.use")
	if !Use {
		return
	}
	var err error

	esConfig := []elastic.ClientOptionFunc{elastic.SetURL(config.GetString("es.url"))}
	if gin.Mode() != gin.ReleaseMode {
		esConfig = append(esConfig, elastic.SetTraceLog(log.New(log.Writer(), "\n", 0)))
	}
	Client, err = elastic.NewClient(esConfig...)
	if err != nil {
		log.Fatal(err)
	}
}

type SetupMethod func() error

func setupMethods() []SetupMethod {
	return []SetupMethod{
		SetupMarkIndex,
	}
}

// Ping and init indices
func Ping() {
	if !Use {
		return
	}

	info, _, err := Client.Ping(config.GetString("es.url")).Do(context.Background())

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
