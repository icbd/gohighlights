package indices

import (
	"context"
	"fmt"
	"github.com/icbd/gohighlights/models"
	"github.com/jinzhu/copier"
	"github.com/olivere/elastic/v7"
	"github.com/spf13/cast"
	"time"
)

const MarkIndexName = "mark"
const MarkIndexMapping = `
{
  "settings": {
    "number_of_shards": 1,
    "number_of_replicas": 0
  },
  "mappings": {
    "properties": {
      "id": {
        "type": "long"
      },
      "created_at": {
        "type": "date"
      },
      "updated_at": {
        "type": "date"
      },
      "user_id": {
        "type": "long"
      },
      "url": {
        "type": "keyword"
      },
      "tag": {
        "type": "keyword"
      },
      "hash_key": {
        "type": "keyword"
      },
      "selection": {
        "type": "text",
        "analyzer": "ik_max_word",
        "search_analyzer": "ik_smart"
      }
    }
  }
}`

func SetupMarkIndex() error {
	ctx := context.Background()
	exists, err := Client.IndexExists(MarkIndexName).Do(ctx)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}
	createIndex, err := Client.CreateIndex(MarkIndexName).BodyString(MarkIndexMapping).Do(ctx)
	if err != nil {
		return err
	}
	if !createIndex.Acknowledged {
		return fmt.Errorf("create %s index failed", MarkIndexName)
	}
	return nil
}

type Mark struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	UserId    uint   `json:"user_id"`
	URL       string `json:"url"`
	Tag       string `json:"tag"`
	HashKey   string `json:"hash_key"`
	Selection string `json:"selection"`
}

func NewMark(markModel *models.Mark) (*Mark, error) {
	markIndex := &Mark{}
	if err := copier.Copy(&markIndex, markModel); err != nil {
		return nil, err
	}
	markIndex.Selection = markModel.SelectionText()

	return markIndex, nil
}

func (m *Mark) Fresh() (*elastic.IndexResponse, error) {
	return Client.
		Index().
		Index(MarkIndexName).
		Id(cast.ToString(m.ID)).
		BodyJson(m).
		Do(context.Background())
}

func DeleteBy(id uint) (*elastic.DeleteResponse, error) {
	return Client.
		Delete().
		Index(MarkIndexName).
		Id(cast.ToString(id)).
		Do(context.Background())
}
