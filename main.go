package main

import (
	"runtime"

	"github.com/aguzmans/cool-storage/configread"
	"github.com/aguzmans/cool-storage/plugins/glacier"
	"github.com/gin-gonic/gin"
)

func coolStorage(config configread.Profile) {

	glacier.Glacier(config)
	router := gin.Default()
	router.GET("/cool-storage/api/v1/ping", func(c *gin.Context) {
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
	config := configread.ParseYamlConfig("conf/cool-api.yaml")
	coolStorage(config.AwsConfig.Profiles[0])
}
