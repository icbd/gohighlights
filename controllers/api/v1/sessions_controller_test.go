package v1

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/icbd/gohighlights/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func TestSessionsCreate(t *testing.T) {
	assert := assert.New(t)

	u := models.FakeUser()
	body, _ := json.Marshal(gin.H{"email": u.Email, "password": u.Password})

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = httptest.NewRequest("POST", "/api/v1/sessions", bytes.NewReader(body))
	SessionsCreate(ctx)

	assert.Equal(http.StatusCreated, w.Code)

	session := models.Session{}
	_ = json.Unmarshal(w.Body.Bytes(), &session)
	assert.Equal(32, len(session.Token))
}
