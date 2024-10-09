package controllers

import (
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/syahlan1/ecommerce-scrapper.git/database"
	"github.com/syahlan1/ecommerce-scrapper.git/models"
)

func GetProduct(c *fiber.Ctx) error {
	var products []models.Product

	// Ambil semua produk dari database
	if err := database.DB.Find(&products).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch products"})
	}

	// Acak urutan produk menggunakan rand.Shuffle
	rand.Seed(time.Now().UnixNano()) // Inisialisasi seed untuk acakan berdasarkan waktu sekarang
	rand.Shuffle(len(products), func(i, j int) {
		products[i], products[j] = products[j], products[i]
	})

	return c.JSON(products)
}
func GetProductById(c *fiber.Ctx) error {
	// Ambil Product ID dari URL params
	productID := c.Params("id")
	var product models.Product

	// Cari produk berdasarkan ID
	if err := database.DB.First(&product, productID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product Not Found"})
	}

	// Parsing JWT secara manual
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

	// Jika user login, simpan riwayat produk yang dilihat
	if userID != 0 {
		userHistory := models.UserHistory{
			UserID:          int(userID),
			ProductID:       int(product.ID),
			ProductName:     product.Name,
			InteractionType: "view",
			Price:           product.Price,
			Category:        product.Category,
			CreatedAt:       time.Now(),
		}

		// Simpan ke database
		if err := database.DB.Create(&userHistory).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save user history"})
		}
	}

	// Berikan respon produk yang ditemukan tanpa tergantung pada status login
	return c.JSON(fiber.Map{
		"product": product,
	})
}

func RecommendProducts(c *fiber.Ctx) error {
	// Ambil user_id dari context
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	var histories []models.UserHistory
	if err := database.DB.Where("user_id = ?", userID).
		Order("created_at desc").Limit(20).Find(&histories).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch user history"})
	}

	productSearchTerms := make([]string, 0)
	for _, history := range histories {
		if history.ProductName != "" {
			productSearchTerms = append(productSearchTerms, history.ProductName)
		} else if history.Category != "" {
			productSearchTerms = append(productSearchTerms, history.Category)
		}
	}

	// Jika tidak ada istilah pencarian dalam histori, ambil 55 produk secara acak
	if len(productSearchTerms) == 0 {
		var allProducts []models.Product
		if err := database.DB.Limit(55).Find(&allProducts).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch products"})
		}

		// Acak produk
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(allProducts), func(i, j int) {
			allProducts[i], allProducts[j] = allProducts[j], allProducts[i]
		})

		return c.JSON(fiber.Map{
			"recommended_products": allProducts,
		})
	}

	// Buat query LIKE untuk setiap istilah pencarian dari histori user
	query := database.DB
	for _, term := range productSearchTerms {
		query = query.Or("name ILIKE ?", "%"+term+"%").Or("category ILIKE ?", "%"+term+"%")
	}

	// Ambil minimal 55 produk dari hasil pencarian berdasarkan histori user
	var recommendedProducts []models.Product
	if err := query.Limit(55).Find(&recommendedProducts).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch recommended products"})
	}

	// Jika produk kurang dari 55, tambahkan produk secara acak untuk melengkapi
	if len(recommendedProducts) < 55 {
		var additionalProducts []models.Product
		neededProducts := 55 - len(recommendedProducts)
		if err := database.DB.Limit(neededProducts).Find(&additionalProducts).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch additional products"})
		}
		recommendedProducts = append(recommendedProducts, additionalProducts...)
	}

	// Acak urutan produk yang direkomendasikan
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(recommendedProducts), func(i, j int) {
		recommendedProducts[i], recommendedProducts[j] = recommendedProducts[j], recommendedProducts[i]
	})

	return c.JSON(fiber.Map{
		"recommended_products": recommendedProducts,
	})
}

func GetLastViewedProducts(c *fiber.Ctx) error {
	// Ambil user_id dari context
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	// Ambil 5 histori terakhir berdasarkan ProductID dari tabel UserHistory
	var history []models.UserHistory
	if err := database.DB.Where("user_id = ?", userID).
		Order("created_at desc").Limit(4).Find(&history).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch user history"})
	}

	// Jika tidak ada histori, kembalikan respon kosong
	if len(history) == 0 {
		return c.JSON(fiber.Map{
			"last_viewed_products": []models.Product{},
		})
	}

	// Ambil ProductID dari histori
	productIDs := make([]int, len(history))
	for i, h := range history {
		productIDs[i] = h.ProductID
	}

	// Ambil produk dari tabel Product berdasarkan ProductID yang didapatkan dari histori
	var products []models.Product
	if err := database.DB.Where("id IN ?", productIDs).Find(&products).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch products"})
	}

	// Return maksimal 5 produk yang terakhir dilihat oleh user
	return c.JSON(fiber.Map{
		"last_viewed_products": products,
	})
}
