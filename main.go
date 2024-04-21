package main

import (
	routes "tag_combination_api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {


	// notion.UpdateDatabase("62d01a2d65284e2f847809abfe3b88da")

  app := fiber.New()

  routes.Register(app)

  app.Listen(":3000")
}
