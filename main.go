package main

import (
	"os"
	"os/exec"
	"tag_combination_api/app/config"
	"tag_combination_api/app/routes"

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

  app := fiber.New()

	buildCmd := exec.Command("npm", "run", "build", "--prefix", "frontend/")
	buildCmd.Run()
	buildCmd.Wait()

	app.Static("/", "./frontend/dist/")
	
	config.ConfigEnv(app)
	config.ConfigDB(app)
	
	// notion.UpdateDatabase("62d01a2d65284e2f847809abfe3b88da")


  routes.Register(app)

  // app.Listen(getPort())

}
