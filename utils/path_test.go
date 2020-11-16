package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func TestRootPath(t *testing.T) {
	assert := assert.New(t)
	assert.Contains(RootPath(), RootName)
}
