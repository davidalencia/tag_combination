package config

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/joho/godotenv"
)

type Product struct {
  gorm.Model
  Code  string
  Price uint
}

type User struct {
  gorm.Model
  Email  string
  PasswordHash string
}


func ConfigDB(app *fiber.App) {
	godotenv.Load("../../.env")
	myEnv, _ := godotenv.Read()
	fmt.Println(myEnv["DSN"])
	db, err := gorm.Open(postgres.Open(myEnv["DSN"]), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

  if err != nil {
    panic("failed to connect database")
  }

  // Migrate the schema
  db.AutoMigrate(&Product{})
  db.AutoMigrate(&User{})

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

}