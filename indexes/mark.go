package indexes

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/icbd/gohighlights/models"
	"github.com/jinzhu/copier"
	"github.com/spf13/cast"
	"time"
)

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

func (m *Mark) Fresh() error {
	requestBody, err := json.Marshal(m)
	if err != nil {
		return err
	}

	req := esapi.IndexRequest{
		Index:      "marks",
		DocumentID: cast.ToString(m.ID),
		Body:       bytes.NewReader(requestBody),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), Client)
	defer res.Body.Close()
	return err
}
