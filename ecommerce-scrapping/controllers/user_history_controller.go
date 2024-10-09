package controllers

import (
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/syahlan1/ecommerce-scrapper.git/database"
	"github.com/syahlan1/ecommerce-scrapper.git/models"
)

func GetStory(c *fiber.Ctx) error {
	var userHistory []models.UserHistory
	database.DB.Find(&userHistory)

	return c.JSON(userHistory)
}

func GetFromSearch(c *fiber.Ctx) error {
	// Buat instance baru dari UserHistory
	userHistory := new(models.UserHistory)

	// Parse JSON input dari body request
	if err := c.BodyParser(userHistory); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Input"})
	}

	// Parsing JWT secara manual untuk memeriksa apakah user login
	cookie := c.Cookies("jwt")
	var userID uint
	if cookie != "" {
		// Parse token JWT
		token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})

		if err == nil && token.Valid {
			// Ambil claims dari token jika valid
			claims, ok := token.Claims.(jwt.MapClaims)
			if ok && token.Valid {
				// Ambil user_id dari JWT claims
				userIdFloat, ok := claims["user_id"].(float64)
				if ok {
					userID = uint(userIdFloat)
				}
			}
		}
	}

	// Simpan riwayat pencarian jika user login
	if userID != 0 {
		userHistory.UserID = int(userID)
		userHistory.InteractionType = "search"
		userHistory.CreatedAt = time.Now()

		// Simpan ke database
		if err := database.DB.Create(&userHistory).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save search history"})
		}
	} else {
		fmt.Println("User not logged in, skipping saving search history.")
	}

	// Pecah input product_name menjadi beberapa kata untuk mencocokkan lebih fleksibel
	searchTerms := strings.Fields(userHistory.ProductName)
	query := database.DB

	// Untuk setiap kata di searchTerms, tambahkan kondisi LIKE ke query
	for _, term := range searchTerms {
		query = query.Or("name ILIKE ?", "%"+term+"%")
	}

	// Lakukan pencarian produk berdasarkan query
	var products []models.Product
	if err := query.Find(&products).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to search products"})
	}

	// Return hasil pencarian, tanpa tergantung pada status login
	return c.JSON(fiber.Map{
		"search_results": products,
		"search_history": func() interface{} {
			if userID != 0 {
				return userHistory
			}
			return nil
		}(),
	})
}
