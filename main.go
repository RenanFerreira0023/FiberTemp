package main

import (
	"fmt"
	"log"
	"net/http"

	"rds_api_2/routers"
)

func main() {
	fmt.Println("Starting server on port 8080...")

	router := routers.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
