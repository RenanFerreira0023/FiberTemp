package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env") // Carrega as variáveis do arquivo .env
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env:", err)
	}
	portSystem := os.Getenv("PORT_SYSTEM")
	fmt.Println("Starting server on port " + portSystem + "...")

	//	router := routers.NewRouter()
	router := fiber.New()
	fmt.Println("Entrou 1")
	router.Listen(":" + portSystem)
}
