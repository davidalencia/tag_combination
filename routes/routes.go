package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)


func Register(app *fiber.App) {
	app.Get("/svg", func(c *fiber.Ctx) error {
		color :=  c.Query("bg")
		if color == "" {
			color = "ffffff"
		}
		textColor := c.Query("textcolor")
		if textColor == "" {
			textColor = "000000"
		}
		c.Set(fiber.HeaderContentType, "image/svg+xml")
		return c.SendString(svg("#"+color, "#"+textColor, c.Query("text")))

	})
}


func svg(color, textColor, text string) string {
	return fmt.Sprintf(`<svg width="1500" height="600" xmlns="http://www.w3.org/2000/svg">
	<rect x="0" y="0" width="1500" height="600" fill="%s" />
	<text fill="%s" font-size="45" x="50%%" y="50%%" dominant-baseline="middle" text-anchor="middle">%s</text>  
 </svg>`, color, textColor, text)
}

// salida a evento