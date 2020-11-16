package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/icbd/gohighlights/models"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const CurrentUser = "CurrentUser"

type BearerVO struct {
	Bearer string `json:"bearer"`
}

func CurrentUserMiddleware(c *gin.Context) {
	data, _ := c.GetRawData()

	bearer := ""
	splits := strings.Split(c.Request.Header.Get("Authorization"), " ")
	if len(splits) == 2 && splits[0] == "Bearer" && len(splits[1]) > 0 {
		bearer = splits[1]
	}

	if bearer == "" {
		bearer = c.Query("bearer")
	}

	if bearer == "" {
		b := BearerVO{}
		log.Println(string(data))
		if err := json.Unmarshal(data, &b); err == nil {
			bearer = b.Bearer
		}
	}

	if bearer != "" {
		s := models.Session{Token: bearer}
		if err := s.FindByToken(); err == nil {
			c.Set(CurrentUser, &s.User)
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
			c.Next()
			return
		}
	}
	c.AbortWithStatus(http.StatusUnauthorized)
	c.Abort()
}
