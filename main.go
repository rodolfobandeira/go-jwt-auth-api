package main

import (
	"crypto/rand"
	"crypto/rsa"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	jwtware "github.com/gofiber/jwt"
)

var (
	// Dynamically creating a new private key on each bootload.
	privateKey *rsa.PrivateKey
)

func main() {
	app := fiber.New()

	rng := rand.Reader
	var err error
	privateKey, err = rsa.GenerateKey(rng, 2048)
	if err != nil {
		log.Fatalf("rsa.GenerateKey: %v", err)
	}

	// Login route
	app.Post("/login", login)

	// Unauthenticated route
	app.Get("/", accessible)

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningMethod: "RS256",
		SigningKey:    privateKey.Public(),
	}))

	// Restricted Routes
	app.Get("/restricted", restricted)

	app.Listen(3000)
}

func login(c *fiber.Ctx) {
	user := c.FormValue("user")
	pass := c.FormValue("pass")

	// curl --location --request POST 'http://localhost:3000/login' \
	// --form 'user="rodolfo"' \
	// --form 'pass="bandeira"'

	// Debug: log.Println(user, pass)
	if user != "rodolfo" || pass != "bandeira" {
		c.SendStatus(fiber.StatusUnauthorized)
		return
	}

	token := jwt.New(jwt.SigningMethodRS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "Rodolfo Bandeira"
	claims["details"] = "Any detail for the user can be added here. :)"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Minute * 2).Unix() // Token is valid for 5 minutes only
	// claims["exp"] = time.Now().Add(time.Hour * (24 * 7)).Unix() // 7 days

	// Generate encoded token and send it as response.
	t, err := token.SignedString(privateKey)
	if err != nil {
		log.Printf("token.SignedString: %v", err)
		c.SendStatus(fiber.StatusInternalServerError)
		return
	}

	c.JSON(fiber.Map{"token": t})
}

func accessible(c *fiber.Ctx) {
	c.Send("Accessible")
}

func restricted(c *fiber.Ctx) {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	details := claims["details"].(string)
	c.Send("Welcome " + name + " \n" + details)
}
