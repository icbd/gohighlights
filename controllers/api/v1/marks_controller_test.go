package v1

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/icbd/gohighlights/controllers/middleware"
	"github.com/icbd/gohighlights/indices"
	"github.com/icbd/gohighlights/models"
	"github.com/icbd/gohighlights/utils"
	"github.com/spf13/cast"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

var currentSession *models.Session

func init() {
	gin.SetMode(gin.TestMode)
	indices.Enable = false

	currentSession, _ = models.FakeUser().GenerateSession()
}

func TestMarksCreate(t *testing.T) {
	assert := assert.New(t)

	params := gin.H{
		"url":       "http://localhost/",
		"tag":       "blue",
		"hash_key":  utils.GenerateToken(10),
		"selection": models.FakeSelection(),
	}
	body, _ := json.Marshal(params)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Set(middleware.CurrentUser, currentSession.User)
	ctx.Request = httptest.NewRequest("POST", "/api/v1/marks", bytes.NewReader(body))
	MarksCreate(ctx)

	assert.Equal(http.StatusCreated, w.Code)

	var resp map[string]interface{}
	_ = json.Unmarshal(w.Body.Bytes(), &resp)
	assert.True(cast.ToUint(resp["id"]) != 0)
}

func TestMarksUpdate(t *testing.T) {
	assert := assert.New(t)

	oldMark := models.FakeMark(currentSession.UserID)

	params := gin.H{"tag": "DIY"}
	body, _ := json.Marshal(params)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Set(middleware.CurrentUser, currentSession.User)
	ctx.Request = httptest.NewRequest("PATCH", "/api/vi/marks/:hash_key", bytes.NewReader(body))
	ctx.Params = gin.Params{gin.Param{Key: "hash_key", Value: oldMark.HashKey}}
	MarksUpdate(ctx)

	assert.Equal(http.StatusOK, w.Code)

	var resp map[string]interface{}
	_ = json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal("DIY", cast.ToString(resp["tag"]))
}

func TestMarksDestroy(t *testing.T) {
	assert := assert.New(t)

	oldMark := models.FakeMark(currentSession.UserID)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Set(middleware.CurrentUser, currentSession.User)
	ctx.Request = httptest.NewRequest("DELETE", "/api/vi/marks/:hash_key", nil)
	ctx.Params = gin.Params{gin.Param{Key: "hash_key", Value: oldMark.HashKey}}
	MarksDestroy(ctx)

	assert.Equal(http.StatusNoContent, w.Code)

	_, err := models.MarkFind(oldMark.ID)
	assert.True(errors.Is(err, gorm.ErrRecordNotFound))
}

func TestMarksQuery(t *testing.T) {
	assert := assert.New(t)

	oldMark := models.FakeMark(currentSession.UserID)

	escapedBase64URL := url.QueryEscape(base64.StdEncoding.EncodeToString([]byte(oldMark.URL)))

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Set(middleware.CurrentUser, currentSession.User)
	ctx.Request = httptest.NewRequest("GET", "/api/vi/marks/query?url="+escapedBase64URL, nil)
	ctx.Params = gin.Params{gin.Param{Key: "hash_key", Value: oldMark.HashKey}}
	MarksQuery(ctx)

	assert.Equal(http.StatusOK, w.Code)

	var resp [](map[string]interface{})
	_ = json.Unmarshal(w.Body.Bytes(), &resp)
	assert.True(len(resp) > 0)

	hash_keys := make(map[string]uint)
	for _, item := range resp {
		hash_keys[cast.ToString(item["hash_key"])] = cast.ToUint(item["id"])
	}
	_, ok := hash_keys[oldMark.HashKey]
	assert.True(ok)
}

func TestMarksSearch(t *testing.T) {
	assert := assert.New(t)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Set(middleware.CurrentUser, currentSession.User)
	ctx.Request = httptest.NewRequest("GET", "/api/vi/marks/search?q=xxx", nil)
	MarksSearch(ctx)

	assert.Equal(http.StatusOK, w.Code)

	var resp [](map[string]interface{})
	_ = json.Unmarshal(w.Body.Bytes(), &resp)
	assert.True(resp != nil)
}
