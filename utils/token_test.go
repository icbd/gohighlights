package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func TestGenerateToken(t *testing.T) {
	assert := assert.New(t)

	t1 := GenerateToken(10)
	t2 := GenerateToken(10)
	assert.NotEqual(t1, t2)
	assert.Equal(20, len(t1))
}
