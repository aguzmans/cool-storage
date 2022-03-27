package main

import (
	// "main/plugins/glacier"
	"runtime"

	"github.com/aguzmans/cool-storage/plugins/glacier"
	"github.com/gin-gonic/gin"
)

func coolStorage() {
	// x := new.Glacier()
	glacier.Glacier()
	router := gin.Default()
	router.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/os", func(c *gin.Context) {
		c.String(200, runtime.GOOS)
	})
	router.Run(":5000")
}

func main() {
	coolStorage()
}
