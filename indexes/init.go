package indexes

import (
	"github.com/elastic/go-elasticsearch/v7"
	"log"
)

var Client *elasticsearch.Client

func init() {
	var err error
	Client, err = elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatal(err)
	}
}
