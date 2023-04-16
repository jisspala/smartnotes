package main

import (
	"fmt"
	"smartnotes/repository"
	"smartnotes/routes"
)

func main() {
	repository.InitDb()
	fmt.Println("Welcome to smartnotes ")
	router := routes.SetupRouter()
	router.Run("localhost:3000")
}
