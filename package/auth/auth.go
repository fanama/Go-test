package auth

import (
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

//User ...
type User struct {
	Username string `json:"user"`
	Password string `json:"pass"`
}

//Login ...
func Login(c *fiber.Ctx) error {
	body := c.Body()

	var user User

	err := json.Unmarshal(body, &user)
	if err != nil {
		return err
	}

	// Throws Unauthorized error
	if user.Password != "password" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = user.Username
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}
