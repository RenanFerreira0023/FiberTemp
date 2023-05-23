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
	err := godotenv.Load(".env") // Carrega as variáveis do arquivo .env
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env:", err)
	}
	portSystem := os.Getenv("PORT_SYSTEM")
	fmt.Println("Starting server on port " + portSystem + "...")

	router := routers.NewRouter()
	fmt.Println("Entrou 1")
	log.Fatal(http.ListenAndServe(":"+portSystem, router))
}
