package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/RenanFerreira0023/FiberTemp/routers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file  ", err.Error())
	}
	port := os.Getenv("PORT_API")
	router := routers.NewRouter()
	fmt.Println("Starting server on port " + port + "...")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
