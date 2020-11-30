package v1

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/icbd/gohighlights/controllers/middleware"
	"github.com/icbd/gohighlights/models"
	"github.com/icbd/gohighlights/utils"
	"github.com/spf13/cast"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCommentsCreate(t *testing.T) {
	assert := assert.New(t)

	oldMark := models.FakeMark(currentSession.UserID)

	params := gin.H{"content": utils.GenerateToken(10)}
	body, _ := json.Marshal(params)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Set(middleware.CurrentUser, currentSession.User)
	ctx.Request = httptest.NewRequest("POST", "/api/v1/marks/:hash_key/comment", bytes.NewReader(body))
	ctx.Params = gin.Params{gin.Param{Key: "hash_key", Value: oldMark.HashKey}}
	CommentsCreate(ctx)

	assert.Equal(http.StatusCreated, w.Code)

	var resp map[string]interface{}
	_ = json.Unmarshal(w.Body.Bytes(), &resp)
	assert.True(cast.ToUint(resp["id"]) != 0)
}

func TestCommentsUpdate(t *testing.T) {
	assert := assert.New(t)

	oldMark := models.FakeMark(currentSession.UserID)
	oldComment := models.FakeComment(currentSession.UserID, oldMark.ID)

	params := gin.H{"content": utils.GenerateToken(10)}
	body, _ := json.Marshal(params)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Set(middleware.CurrentUser, currentSession.User)
	ctx.Request = httptest.NewRequest("PATCH", "/api/vi/marks/:hash_key/comment", bytes.NewReader(body))
	ctx.Params = gin.Params{gin.Param{Key: "hash_key", Value: oldMark.HashKey}}
	CommentsUpdate(ctx)

	assert.Equal(http.StatusOK, w.Code)

	var resp map[string]interface{}
	_ = json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(params["content"], cast.ToString(resp["content"]))
	assert.Equal(oldComment.ID, cast.ToUint(resp["id"]))
}

func TestCommentsDestroy(t *testing.T) {
	assert := assert.New(t)

	oldMark := models.FakeMark(currentSession.UserID)
	models.FakeComment(currentSession.UserID, oldMark.ID)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Set(middleware.CurrentUser, currentSession.User)
	ctx.Request = httptest.NewRequest("DELETE", "/api/vi/marks/:hash_key/comment", nil)
	ctx.Params = gin.Params{gin.Param{Key: "hash_key", Value: oldMark.HashKey}}
	CommentsDestroy(ctx)

	assert.Equal(http.StatusNoContent, w.Code)
}
