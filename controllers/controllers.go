package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vincentxvega/birthdays/models"
)

// Get all birthdays
func GetBirthdays(c *gin.Context) {
	var birthdays []models.Birthday
	models.DB.Find(&birthdays)

	c.JSON(http.StatusOK, gin.H{"data": birthdays})
}
