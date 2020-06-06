package services

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitServer() {
	r := gin.Default()
	r.Use(checkIfAuthorized)
	r.GET("/ping", func(c *gin.Context) {
		c.Request.Header.Get("")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}

func checkIfAuthorized(c *gin.Context) {
	uA := c.Request.Header.Get("Authorization")
	fmt.Println(uA)
}
