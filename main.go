package main

import (
	"fmt"
	"smartnotes/docs"
	"smartnotes/repository"
	"smartnotes/routes"
)

// @title SmartNotes API
// @version 1.0
// @description This is a Simple SmartNotes API server.
// @host localhost:8001
// @BasePath /
// @query.collection.format multi
func main() {

	// os.Setenv("POSTGRES_USER", "user")
	// os.Setenv("POSTGRES_PASSWORD", "admin")
	// os.Setenv("POSTGRES_PORT", "5432")
	// os.Setenv("POSTGRES_HOST", "localhost")
	// os.Setenv("POSTGRES_DB", "notepad")

	repository.InitDb()
	fmt.Println("Welcome to smartnotes ")
	router := routes.SetupRouter()
	router.Run(":8001")
	docs.SwaggerInfo.Title = "Swagger Example API"

}
