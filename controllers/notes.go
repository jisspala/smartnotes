package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Note struct {
	ID    int
	Title string
	Text  string
}

var notes = []Note{{
	ID:    1,
	Title: "First Note",
	Text:  "Welcome to smartnotes",
}}

func GetNotes(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, notes)
}
