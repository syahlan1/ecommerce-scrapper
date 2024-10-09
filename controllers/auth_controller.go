package controllers

import (
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/syahlan1/ecommerce-scrapper.git/database"
	"github.com/syahlan1/ecommerce-scrapper.git/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

)

const SecretKey = "secret" // Ganti dengan kunci rahasia yang lebih aman

// Register User
func Register(c *fiber.Ctx) error {
	data := new(models.User)

	// Parse input JSON dari request
	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// Cek apakah username sudah ada
	var existingUser models.User
	err := database.DB.Where("username = ?", data.Username).First(&existingUser).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		// Ini adalah error lain selain record not found, misalnya error dari koneksi database
		fmt.Println("Error when checking username:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error checking username"})
	}
	if err == nil {
		// Jika tidak ada error, artinya record ditemukan dan username sudah ada
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Username already exists"})
	}

	// Cek apakah email sudah ada
	err = database.DB.Where("email = ?", data.Email).First(&existingUser).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		// Ini adalah error lain selain record not found, misalnya error dari koneksi database
		fmt.Println("Error when checking email:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error checking email"})
	}
	if err == nil {
		// Jika tidak ada error, artinya record ditemukan dan email sudah ada
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Email already exists"})
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
	}
	data.Password = string(hashedPassword)

	// Simpan user baru ke database
	if err := database.DB.Create(data).Error; err != nil {
		fmt.Println("Error saving user:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
	}

	return c.JSON(fiber.Map{"message": "User registered successfully"})
}

// Login User
func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// Cari user berdasarkan email
	var user models.User
	if err := database.DB.Where("username = ?", data["username"]).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	// Cek apakah password cocok dengan password yang di-hash di database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Incorrect password"})
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not login"})
	}

	// Simpan token ke cookies
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 72),
		HTTPOnly: true,  // Jangan biarkan JavaScript membaca cookie
		Secure:   false, // Ubah ke true jika di production dengan HTTPS
		SameSite: "none",
	}

	c.Cookie(&cookie)

	// Update status is_login
	user.IsLogin = 1
	if err := database.DB.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update user status"})
	}

	return c.JSON(fiber.Map{"message": "Login successful"})
}

func GetUser(c *fiber.Ctx) error {
	// Ambil data user dari middleware
	user := c.Locals("user").(models.User)
	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	// Ambil JWT dari cookie
	cookie := c.Cookies("jwt")

	// Parse token
	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	// Ambil user_id dari claims
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(float64)

	// Temukan user dan update status is_login
	var user models.User
	database.DB.First(&user, int(userId))
	user.IsLogin = 0
	database.DB.Save(&user)

	// Hapus cookie
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	})

	return c.JSON(fiber.Map{"message": "Logout successful"})
}
