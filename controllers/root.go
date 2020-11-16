package controllers

import (
	"github.com/gin-gonic/gin"
	"time"
)

// root and health check
func Root(c *gin.Context) {
	c.JSON(200, gin.H{
		"health": "ok",
		"now":    time.Now().Format(time.RFC3339),
	})
}
