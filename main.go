package main

import (
	//	"fmt"
	//	"log"
	//	"net/http"
	//	"os"

	//	"github.com/RenanFerreira0023/FiberTemp/routers"
	"github.com/gofiber/fiber/v2"
	//	"github.com/joho/godotenv"
)

/*
	func main2() {
		err := godotenv.Load(".env") // Carrega as variáveis do arquivo .env
		if err != nil {
			log.Fatal("Erro ao carregar o arquivo .env:", err)
		}
		portSystem := os.Getenv("PORT_SYSTEM")
		fmt.Println("Starting server on port " + portSystem + "...")

		router := routers.NewRouter()

		log.Fatal(http.ListenAndServe(":"+portSystem, router))
	}
*/
func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
