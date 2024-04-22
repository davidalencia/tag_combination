package routes

import (
	"fmt"
	"tag_combination_api/app/config"
	"time"

	"github.com/go-playground/validator/v10"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)


func Register(app *fiber.App) {

	jwtHandler := jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	})
	
	app.Get("/api/svg", svgHandler)
	
	app.Post("/api/login", loginHandler)
	app.Post("/api/register", registerHandler)
	

	// ------ protected ---------------
	app.Post("/api/cover", jwtHandler, coverHandler)
}

//------- helpers --------------------------
var validate = validator.New(validator.WithRequiredStructEnabled())
func parseAndValidate[T any](c *fiber.Ctx) (*T, error){
	body := new(T)
	if err := c.BodyParser(body); err != nil {
		return nil, err
	}
	if err := validate.Struct(body); err != nil {
		c.SendString(fmt.Sprint(err))
		return nil, err
	}
	return body, nil
}

func svg(color, textColor, text string) string {
	return fmt.Sprintf(`<svg width="1500" height="620" xmlns="http://www.w3.org/2000/svg">
	<rect x="0" y="0" width="1500" height="620" fill="%s" />
	<text fill="%s" font-size="100" x="50%%" y="50%%" dominant-baseline="middle" text-anchor="middle">%s</text>  
 </svg>`, color, textColor, text)
}


//------- handlers ---------------------------

func registerHandler(c *fiber.Ctx) error {
	body, err := parseAndValidate[struct {
		User string `json:"user" validate:"required,email"`
		Pass string `json:"pass" validate:"required"`
	}](c)
	if err!=nil {
		return c.SendString(fmt.Sprint(err))
	}

	db := c.Locals("db").(*gorm.DB)
	password, _ := bcrypt.GenerateFromPassword([]byte(body.Pass), 14)
	user := config.User{Email: body.User, PasswordHash: string(password)}
	db.Create(&user) 
	return c.JSON(fiber.Map{"email": user.Email})
}


func loginHandler(c *fiber.Ctx) error {
	body, err := parseAndValidate[struct {
		User string `json:"user" validate:"required,email"`
		Pass string `json:"pass" validate:"required"`
	}](c)
	if err!=nil {
		return c.SendString(fmt.Sprint(err))
	}

	db := c.Locals("db").(*gorm.DB)

	var user config.User
	db.First(&user, "email = ?", body.User) 
	failed := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(body.Pass))
	if failed != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"user":  body.User,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	//todo: cambiar env
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}

func svgHandler(c *fiber.Ctx) error{
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

}

func coverHandler(c *fiber.Ctx) error {

	return c.SendString("test")
}