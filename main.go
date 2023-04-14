package main

import (
	"fmt"
	"smartnotes/routes"
)

func main() {
	fmt.Println("Welcome to smartnotes ")
	router := routes.SetupRouter()
	router.Run("localhost:3000")
}
