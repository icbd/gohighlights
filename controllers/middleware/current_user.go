package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/icbd/gohighlights/models"
	"net/http"
	"strings"
)

const CurrentUser = "CurrentUser"

type BearerVO struct {
	Bearer string `json:"bearer"`
}

func CurrentUserMiddleware(c *gin.Context) {
	bearer := ""
	splits := strings.Split(c.Request.Header.Get("Authorization"), " ")
	if len(splits) == 2 && splits[0] == "Bearer" && len(splits[1]) > 0 {
		bearer = splits[1]
	}

	//// query params
	//if bearer == "" {
	//	bearer = c.Query("bearer")
	//}
	//
	//// request body
	//if bearer == "" && c.Request.Body != nil {
	//	if data, err := ioutil.ReadAll(c.Request.Body); err == nil {
	//		b := BearerVO{}
	//		if err := json.Unmarshal(data, &b); err == nil {
	//			bearer = b.Bearer
	//		}
	//		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
	//	}
	//}

	if bearer != "" {
		s := models.Session{Token: bearer}
		if err := s.FindByToken(); err == nil {
			c.Set(CurrentUser, &s.User)
			c.Next()
			return
		}
	}
	c.AbortWithStatus(http.StatusUnauthorized)
	c.Abort()
}
