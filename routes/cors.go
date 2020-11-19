package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORS(r *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowBrowserExtensions = true
	config.AllowOrigins = []string{"chrome-extension://homlcfpinafhealhlmjkmdjdejppmmlk"}
	r.Use(cors.New(config))
}
