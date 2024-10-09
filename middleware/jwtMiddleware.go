package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/syahlan1/ecommerce-scrapper.git/database"
	"github.com/syahlan1/ecommerce-scrapper.git/models"
)

const SecretKey = "secret"

func AuthRequired(c *fiber.Ctx) error {
	// Ambil JWT dari cookie
	cookie := c.Cookies("jwt")
	if cookie == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	// Parse token JWT
	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	// Ambil claims dari token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
	}

	// Ambil user_id dari JWT claims
	userIdFloat, ok := claims["user_id"].(float64)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user ID in token"})
	}

	// Konversi userId dari float64 ke uint
	userId := uint(userIdFloat)

	// Cari user berdasarkan user_id dari database
	var user models.User
	if err := database.DB.First(&user, userId).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User not found"})
	}

	// Simpan data user_id dan objek user di context
	c.Locals("user_id", user.ID) // Menyimpan user_id di context
	c.Locals("user", user)       // Menyimpan objek user di context

	// Lanjutkan ke handler berikutnya
	return c.Next()
}
