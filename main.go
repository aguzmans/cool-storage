package main

import (
    "runtime"
    "github.com/gin-gonic/gin"
    "plugins/glacier"
)

func main() {
    Glacier()
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
