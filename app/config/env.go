package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)


func ConfigEnv(app *fiber.App) {
	godotenv.Load("../../.env")
	myEnv, _ := godotenv.Read()

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("env", myEnv)
		return c.Next()
	})

}