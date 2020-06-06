package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/liridonRama/birthday-card-app/src/services"
)

func main() {
	services.InitServer()
}

func CheckIfAuthorized(c *gin.Context) {
	fmt.Println("password checked")
}
