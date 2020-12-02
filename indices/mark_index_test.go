package indices

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/icbd/gohighlights/models"
	"github.com/olivere/elastic/v7"
	"github.com/spf13/cast"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func cleanAllIndices() {
	indices, _ := elastic.NewCatIndicesService(Client()).Do(context.Background())
	for _, index := range indices {
		if strings.HasPrefix(index.Index, indexNamePrefix) {
			Client().DeleteIndex(index.Index).Do(context.Background())
		}
	}
}

func indexTotalCount(indexName string) int64 {
	count, _ := Client().Count(indexName).Do(context.Background())
	return count
}

func TestSetupMarkIndex(t *testing.T) {
	assert := assert.New(t)
	cleanAllIndices()
	Client().DeleteIndex(IndexName(MarkIndexName)).Do(context.Background())

	if err := SetupMarkIndex(); err != nil {
		assert.Error(err)
	}
	exist, _ := Client().IndexExists(IndexName(MarkIndexName)).Do(context.Background())
	assert.True(exist)
}

func TestMarkIndex_Fresh(t *testing.T) {
	assert := assert.New(t)
	cleanAllIndices()

	mark := models.FakeMark(1)
	resp, err := NewMarkIndex(mark).Fresh()
	assert.Nil(err)
	assert.Equal(cast.ToString(mark.ID), resp.Id)

	time.Sleep(time.Second)
	assert.Equal(cast.ToInt64(1), indexTotalCount(IndexName(MarkIndexName)))
}

func TestMarkIndex_MarkIndexDelete(t *testing.T) {
	assert := assert.New(t)
	cleanAllIndices()

	mark := models.FakeMark(1)
	NewMarkIndex(mark).Fresh()
	_, err := MarkIndexDelete(mark.ID)
	assert.Nil(err)

	time.Sleep(time.Second)
	assert.Equal(cast.ToInt64(0), indexTotalCount(IndexName(MarkIndexName)))
}
