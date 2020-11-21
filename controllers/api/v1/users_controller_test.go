package v1

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/icbd/gohighlights/controllers/middleware"
	"github.com/icbd/gohighlights/models"
	"github.com/spf13/cast"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func TestUsersCreate(t *testing.T) {
	assert := assert.New(t)

	u := models.FakeUser()
	body, _ := json.Marshal(gin.H{"email": u.Email, "password": u.Password})
	models.DB().Unscoped().Delete(&u)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = httptest.NewRequest("POST", "/api/v1/users", bytes.NewReader(body))
	ctx.Request.Header.Set("Content-Type", "application/json")

	UsersCreate(ctx)
	assert.Equal(http.StatusCreated, w.Code)
	var resp map[string]interface{}
	_ = json.Unmarshal(w.Body.Bytes(), &resp)
	assert.True(cast.ToUint(resp["id"]) != 0)
}

func TestUsersCreate_BadRequest(t *testing.T) {
	assert := assert.New(t)

	u := models.FakeUser() // already created by u.Email
	body, _ := json.Marshal(gin.H{"email": u.Email, "password": u.Password})

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = httptest.NewRequest("POST", "/api/v1/users", bytes.NewReader(body))

	UsersCreate(ctx)
	assert.Equal(http.StatusBadRequest, w.Code)
}

// Controller test won't go through the middleware
func TestUsersShow_NotFound(t *testing.T) {
	assert := assert.New(t)

	u := models.FakeUser()

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = httptest.NewRequest("GET", "/api/v1/users/"+cast.ToString(u.ID), nil)
	ctx.Set(middleware.CurrentUser, u)

	UsersShow(ctx)
	assert.Equal(http.StatusOK, w.Code)
}
