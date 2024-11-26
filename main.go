package main

import (
	"log"

	"github.com/danigilang17/be-simwas/services"

	"github.com/danigilang17/be-simwas/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	services.InitRedis()

	app := fiber.New()

	app.Post("/login", handlers.Login)
	app.Post("/verify", handlers.VerifyOTP)

	log.Fatal(app.Listen(":3000"))
}
