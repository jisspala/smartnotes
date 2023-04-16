package controllers

import (
	"net/http"
	"smartnotes/models"
	"smartnotes/repository"
	"smartnotes/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetNotes(c *gin.Context) {
	var allNotes []models.Note

	var db *gorm.DB = repository.GetDB()

	db.Find(&allNotes)

	c.JSON(http.StatusOK, allNotes)
}

func AddNote(c *gin.Context) {
	var newNote models.Note
	var db = repository.GetDB()

	err := c.BindJSON(&newNote)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{Message: "Invalid request body"})
		return
	}

	if newNote.Title == "" || newNote.Text == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{Message: "Title and Text are required"})
		return
	}

	result := db.Create(&newNote) // pass pointer of data to Create

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{Message: "Not able to create note"})
		return
	}
	c.JSON(http.StatusCreated, newNote)
}
