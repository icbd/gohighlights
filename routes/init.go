package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var r = gin.New()

func init() {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())          // CORS: Allow All
	r.MaxMultipartMemory = 1 << 20 // 1 MiB

	Root(r)
	ApiV1Public(r)
	ApiV1Auth(r)
}

func R() *gin.Engine {
	return r
}
