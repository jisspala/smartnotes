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

// @Summary get health
// @ID get-health
// @Produce json
// @Success 200 {object} models.Health
// @Router /v1/health [get]
func GetHealth(c *gin.Context) {
	c.JSON(http.StatusOK, models.Health{Status: "OK"})
}

// @Summary get all items in the notes list
// @ID get-all-notes
// @Produce json
// @Success 200 {object} models.Note
// @Router /v1/notes [get]
func GetNotes(c *gin.Context) {
	var allNotes []models.Note

	var db *gorm.DB = repository.GetDB()

	db.Find(&allNotes)

	c.JSON(http.StatusOK, allNotes)
}

// @Summary add a new note to the models.Note list
// @ID create-note
// @Produce json
// @Param data body models.SingleNote true "models.SingleNote data"
// @Success 200 {object} models.SingleNote
// @Router /v1/note [post]
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

// @Summary delete a note item by ID
// @ID delete-note-by-id
// @Produce json
// @Param id path string true "note ID"
// @Success 200 {object} models.Success
// @Router /v1/note/{id} [delete]
func DeleteNote(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var db = repository.GetDB()
	db.Delete(&models.Note{}, id)

	// result:=db.Debug().First(&models.Note{}, id)
	result := db.Debug().Delete(&models.Note{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

// @Summary delete notes by IDs
// @ID delete-notes-by-ids
// @Produce json
// @Param data body models.MultiID[] true "models.MultiID[] data"
// @Success 200 {object} models.Success
// @Router /v1/note [delete]
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

// @Summary update a note by ID
// @ID update-note-by-id
// @Produce json
// @Param id path string true "note ID"
// @Param data body models.SingleNote true "models.SingleNote data"
// @Success 200 {object} models.SingleNote
// @Success 200 {object} models.Success
// @Router /v1/note/{id} [put]
func UpdateNote(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var note models.Note
	var isNote models.Note
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

	if err := db.Debug().First(&isNote, id).Error; err != nil {
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorResponse{Message: "Note not found"})
		}
		return
	}

	result := db.Debug().Save(&models.Note{ID: id, Title: note.Title, Text: note.Text})

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "no matching records found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}
