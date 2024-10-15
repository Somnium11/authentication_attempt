package main

import (
    "github.com/gin-gonic/gin"
    "github.com/spf13/viper"
    "github.com/golang-jwt/jwt/v5"

)

func main() {
    r := gin.Default()

    r.GET("/hello", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Hello, world!",
        })
    })

    r.Run()
}


