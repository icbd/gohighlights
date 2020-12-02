package indices

import (
	"context"
	"github.com/icbd/gohighlights/models"
	"github.com/olivere/elastic/v7"
	"github.com/spf13/cast"
)

// CommentIndex is part of MarkIndex.
// Only update CommentIndex on existed MarkIndex.
type CommentIndex struct {
	err     error
	indexID uint

	Comment string `json:"comment"`
}

func NewCommentIndex(c *models.Comment) *CommentIndex {
	return &CommentIndex{Comment: c.Content, indexID: c.MarkID}
}

/*
POST /mark/_update/:markID
{
  "doc": {
    "comment": "actually"
  }
}
*/
func (commentIndex *CommentIndex) Update() (*elastic.UpdateResponse, error) {
	if !Enable {
		return nil, NotEnabledError
	}
	if commentIndex.err != nil {
		return nil, commentIndex.err
	}

	return Client().
		Update().
		Index(IndexName(MarkIndexName)).
		Id(cast.ToString(commentIndex.indexID)).
		Doc(commentIndex).
		Do(context.Background())
}
