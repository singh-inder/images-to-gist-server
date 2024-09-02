package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/images-to-gist-server/handlers"
)

func main() {
	PORT := os.Getenv("PORT")

	if len(PORT) == 0 {
		PORT = ":5000"
	}

	if !strings.HasPrefix(PORT, ":") {
		PORT = ":" + PORT
	}

	app := fiber.New()
	app.Use(helmet.New())

	app.Get("/", handlers.ServeImage)

	fmt.Printf("Starting server on PORT %s\n", PORT)

	if err := app.Listen(PORT); err != nil {
		log.Fatal(err)
	}
}
