package handlers

import (
	"crypto/rand"
	"fmt"
	"net/http"

	"github.com/danigilang17/be-simwas/services"
	"github.com/danigilang17/be-simwas/utils"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	type Request struct {
		Phone string `json:"phone"`
		Role  string `json:"role"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	otp := generateOTP()
	if err := services.SendOTPWhatsApp(req.Phone, otp); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to send OTP"})
	}

	err := services.SetRedisKey(req.Phone, otp, 300)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to store OTP"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "OTP sent"})
}

func VerifyOTP(c *fiber.Ctx) error {
	type Request struct {
		Phone string `json:"phone"`
		OTP   string `json:"otp"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	storedOTP, err := services.GetRedisKey(req.Phone)
	if err != nil || storedOTP != req.OTP {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid OTP"})
	}

	token, err := utils.GenerateJWT("user") // Ganti "user" dengan role dinamis
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"token": token})
}

func generateOTP() string {
	num := make([]byte, 8)
	rand.Read(num)
	n := int(num[0])<<24 | int(num[1])<<16 | int(num[2])<<8 | int(num[3])
	if n < 0 {
		n = -n
	}
	return fmt.Sprintf("%04d", n%10000)
}
