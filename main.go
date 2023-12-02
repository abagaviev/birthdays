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
	router.GET("/birthdays/:id", controllers.GetBirthdayByID)
	router.POST("/birthdays", controllers.CreateBirthday)
	router.PATCH("/birthdays/:id", controllers.UpdateBirthday)
	router.DELETE("/birthdays/:id", controllers.DeleteBirthday)

	router.Run("localhost:8080")
}
