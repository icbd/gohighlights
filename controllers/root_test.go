package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func TestRoot(t *testing.T) {
	assert := assert.New(t)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequest("GET", "/", nil)

	Root(ctx)

	assert.Equal(200, w.Code)

	var body map[string]string
	_ = json.Unmarshal(w.Body.Bytes(), &body)
	assert.Equal("ok", body["health"])
}
