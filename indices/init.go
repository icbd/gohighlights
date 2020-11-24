package indices

import (
	"context"
	"github.com/icbd/gohighlights/config"
	"github.com/olivere/elastic/v7"
	"log"
)

var Client *elastic.Client
var Use bool // if not, jump ES processing

func init() {
	Use = config.GetBool("es.use")
	var err error
	Client, err = elastic.NewClient(
		elastic.SetURL(config.GetString("es.url")),
	)
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
