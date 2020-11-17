package config

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func TestGet(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("mysql", cast.ToString(Get("db.type")))
}

func TestGetString(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("mysql", GetString("db.type"))
}

func TestSet(t *testing.T) {
	assert := assert.New(t)
	configKey := "tempConfigKey"
	configValue := "TEMPCONFIGVALUE"
	Set(configKey, configValue)
	assert.Equal(configValue, GetString(configKey))
}
