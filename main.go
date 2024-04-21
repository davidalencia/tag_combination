package main

import (
	"os"
	routes "tag_combination_api/routes"

	"github.com/gofiber/fiber/v2"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	} else {
		port = ":" + port
	}

	return port
}

func main() {


	// notion.UpdateDatabase("62d01a2d65284e2f847809abfe3b88da")

	

  app := fiber.New()

  routes.Register(app)

  app.Listen(getPort())
}
