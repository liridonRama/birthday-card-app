package main

import (
	"fmt"

	models "github.com/liridonRama/birthday-card-app/models/server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run()

	u := models.User{
		ID:       primitive.NewObjectID(),
		Email:    "liridon.rama.ch@gmail.com",
		Password: "1233",
	}
	u.HashPassword()
	fmt.Println(u)
}
