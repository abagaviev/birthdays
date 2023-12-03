package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vincentxvega/birthdays/models"
)

type CreateBirthdayInput struct {
	Name        string `json:"name" binding:"required"`
	Surname     string `json:"surname" binding:"required"`
	DateOfBirth string `json:"date_of_birth" binding:"required"`
}

type UpdateBirthdayInput struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	DateOfBirth string `json:"date_of_birth"`
}

// Get all birthdays
func GetBirthdays(c *gin.Context) {
	var birthdays []models.Birthday
	models.DB.Find(&birthdays)

	c.JSON(http.StatusOK, gin.H{"data": birthdays})
}

// Get book by id
func GetBirthdayByID(c *gin.Context) {
	var birthday models.Birthday

	if err := models.DB.Where("id = ?", c.Param("id")).First(&birthday).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "birthday not found!"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"data": birthday})
}

// Create birthday

func CreateBirthday(c *gin.Context) {
	// Validate input
	var input CreateBirthdayInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tm, err := time.Parse("2006-01-02", input.DateOfBirth)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format use YYYY-MM-DD format"})
		return
	} else {
		birthday := models.Birthday{Name: input.Name, Surname: input.Surname, DateOfBirth: tm}
		models.DB.Create(&birthday)

		c.JSON(http.StatusOK, gin.H{"data": birthday})
	}

}

// Update birthday by ID
func UpdateBirthday(c *gin.Context) {
	// Get model if exist
	var birthday models.Birthday
	if err := models.DB.Where("id = ?", c.Param("id")).First(&birthday).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "birthday not found!"})
		return
	}

	// Validate input
	var input UpdateBirthdayInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&birthday).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": birthday})
}

func DeleteBirthday(c *gin.Context) {
	// Get model if exist
	var birthday models.Birthday
	if err := models.DB.Where("id = ?", c.Param("id")).First(&birthday).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "birthday not found!"})
		return
	}

	models.DB.Delete(&birthday)

	c.JSON(http.StatusOK, gin.H{"message": "birthday deleted!"})
}
