package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vincentxvega/birthdays/controllers"
	"github.com/vincentxvega/birthdays/models"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase()

	router.GET("/birthdays", controllers.GetBirthdays)

	router.Run("localhost:8080")
}
