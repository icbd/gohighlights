package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/icbd/gohighlights/controllers"
	v1 "github.com/icbd/gohighlights/controllers/api/v1"
	"github.com/icbd/gohighlights/controllers/middleware"
)

func Root(r *gin.Engine) {
	r.GET("/", controllers.Root)
}

func ApiV1Public(r *gin.Engine) {
	rg := r.Group("/api/v1")

	rg.POST("/users", v1.UsersCreate)
	rg.POST("/sessions", v1.SessionsCreate)
}

func ApiV1Auth(r *gin.Engine) {
	rg := r.Group("/api/v1")
	rg.Use(middleware.CurrentUserMiddleware)

	rg.GET("/users/:id", v1.UsersShow)

	rg.POST("/marks", v1.MarksCreate)
	rg.DELETE("/marks/:hash_key", v1.MarksDestroy)
	rg.PATCH("/marks/:hash_key", v1.MarksUpdate)
	rg.GET("/marks/query", v1.MarksQuery)
	rg.GET("/marks/search", v1.MarksSearch)
	rg.GET("/marks", v1.MarksIndex)
}
