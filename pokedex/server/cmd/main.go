package main

import (
	"fmt"
	"pokedex/internals/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Welcome to the Pokédex server!")
	fmt.Println("Starting server...")
	runServer()
}

func runServer() {
	r := gin.Default()
	handlers.InitializeRouter(r)

	if err := r.Run(":8081"); err != nil {
		panic(err)
	}
	fmt.Println("Server stopped.")
}
