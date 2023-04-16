package routes

import (
	"smartnotes/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1") // v1 version api routes

	v1.GET("/notes", controllers.GetNotes)
	v1.POST("/note", controllers.AddNote)

	return router
}
