package indices

import (
	"context"
	"encoding/json"
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
	if !Use {
		return nil
	}
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

type MarkIndex struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	UserID    uint   `json:"user_id"`
	URL       string `json:"url"`
	Tag       string `json:"tag"`
	HashKey   string `json:"hash_key"`
	Selection string `json:"selection"`
}

func NewMark(markModel *models.Mark) (*MarkIndex, error) {
	if !Use {
		return nil, nil
	}
	markIndex := &MarkIndex{}
	if err := copier.Copy(&markIndex, markModel); err != nil {
		return nil, err
	}
	markIndex.Selection = markModel.SelectionText()

	return markIndex, nil
}

func (m *MarkIndex) Fresh() (*elastic.IndexResponse, error) {
	if !Use {
		return nil, nil
	}
	return Client.
		Index().
		Index(MarkIndexName).
		Id(cast.ToString(m.ID)).
		BodyJson(m).
		Do(context.Background())
}

func DeleteBy(id uint) (*elastic.DeleteResponse, error) {
	if !Use {
		return nil, nil
	}
	return Client.
		Delete().
		Index(MarkIndexName).
		Id(cast.ToString(id)).
		Do(context.Background())
}

/**
GET /mark/_search
{
  "query": {
    "bool": {
      "must": [
        {
          "term": {
            "user_id": {
              "value": 1
            }
          }
        },
        {
          "match": {
            "selection": "过滤"
          }
        }
      ]
    }
  },
  "_source": [
    "id"
  ]
}
*/
func Query(u *models.User, text string) (marks []*MarkIndex) {
	marks = make([]*MarkIndex, 0)
	if !Use {
		return marks
	}
	boolQuery := elastic.NewBoolQuery()
	boolQuery.Must(
		elastic.NewTermQuery("user_id", u.ID),
		elastic.NewMatchQuery("selection", text),
	)
	result, err := Client.
		Search().
		Index(MarkIndexName).
		Query(boolQuery).
		Sort("id", false).
		Do(context.Background())
	if err != nil {
		return
	}

	for _, item := range result.Hits.Hits {
		m := MarkIndex{}
		if err := json.Unmarshal(item.Source, &m); err == nil {
			marks = append(marks, &m)
		}
	}
	return marks
}
