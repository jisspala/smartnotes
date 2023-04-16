package controllers

import (
	"fmt"
	"net/http"
	"smartnotes/models"
	"smartnotes/repository"
	"smartnotes/utils"
	"strconv"

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

func DeleteNote(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var db = repository.GetDB()
	db.Delete(&models.Note{}, id)

	result := db.Debug().Delete(&models.Note{}, id)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "no matching records found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

func MultiNoteDelete(c *gin.Context) {
	var ids models.MultiID
	err := c.ShouldBindJSON(&ids)
	fmt.Print((ids), err)
	var db = repository.GetDB()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{Message: "Invalid request body"})
		return
	}
	result := db.Where("id IN (?)", ids.IDs).Delete(&models.Note{})

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "no matching records found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

func UpdateNote(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var note models.Note
	var db = repository.GetDB()
	err := c.BindJSON(&note)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{Message: "Invalid request body"})
		return
	}

	if note.Title == "" || note.Text == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{Message: "Title and Text are required"})
		return
	}

	if err := db.First(&note, id).Error; err != nil {
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{Message: "Note not found"})
		}
		return
	}

	result := db.Save(&models.Note{ID: id, Title: note.Title, Text: note.Text})

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "no matching records found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}
