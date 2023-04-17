package routes

import (
	"smartnotes/controllers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1") // v1 version api routes

	v1.GET("/notes", controllers.GetNotes)
	v1.POST("/note", controllers.AddNote)
	v1.DELETE("/note/:id", controllers.DeleteNote)
	v1.DELETE("/note/", controllers.MultiNoteDelete)
	v1.PUT("/note/:id", controllers.UpdateNote)
	v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
