package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
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

var user *models.User
var session *models.Session

func init() {
	user = models.FakeUser()
	session, _ = user.GenerateSession()
}

func TestCurrentUserMiddleware(t *testing.T) {
	assert := assert.New(t)

	routePath := "/CurrentUserMiddleware"
	r := fakeRouter(routePath)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", routePath, nil)
	req.Header.Set("Authorization", "Bearer "+session.Token)
	r.ServeHTTP(w, req)
	assert.Equal(http.StatusOK, w.Code)

	resp := gin.H{}
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(cast.ToUint(resp["id"]), user.ID)
}

func TestCurrentUserMiddleware_Unauthorized(t *testing.T) {
	assert := assert.New(t)

	routePath := "/CurrentUserMiddleware"
	r := fakeRouter(routePath)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", routePath, nil)
	req.Header.Set("Authorization", "Bearer InvalidToken")
	r.ServeHTTP(w, req)
	assert.Equal(http.StatusUnauthorized, w.Code)
}

func fakeRouter(routePath string) *gin.Engine {
	r := gin.New()
	r.Use(CurrentUserMiddleware)
	r.GET(routePath, func(c *gin.Context) {
		u, _ := c.Get(CurrentUser)
		c.JSON(200, u)
	})
	return r
}
