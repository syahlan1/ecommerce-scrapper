package controllers

// import(
// 	"github.com/gofiber/fiber/v2"
//     "github.com/syahlan1/ecommerce-scrapper.git/database"
//     "github.com/syahlan1/ecommerce-scrapper.git/models"
// )

// func RecommendProduct(c *fiber.Ctx) error {
// 	// get user_id from context
// 	userID, ok := c.Locals("user_id").(uint)
// 	if !ok {
// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
// 	}

// 	// Get limit 20 history
// 	var histories []models.UserHistory
// 	if err := database.DB.Where("user_id = ?", userID).
// 		Order("created_at desc").Limit(20).Find(&histories).Error; err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch user history"})
// 	}

// 	// get product_id from histori and collect recommend products
// 	productIDs := make([]uint, 0)
// 	for _, history := range histories {
// 		if history.ProductID != 0 {
// 			productIDs = append(productIDs, uint(history.ProductID))
// 		}
// 	}

// 	// Get Product Category in history user
// 	var categories []string
// 	if len(productIDs) > 0 {
// 		var products []models.Product
// 		if err := database.DB.Where("id IN (?)", productIDs).Find(&products).Error; err != nil {
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch products from history"})
// 		}

// 		// Collect Category from product in history
// 		for _, product := range products {
// 			categories = append(categories, product.Category)
// 		}
// 	}

// 	// Jika kategori tidak ditemukan, ambil 55 produk secara acak
// 	if len(categories) == 0 {
// 		var allProducts []models.Product
// 		if err := database.DB.Limit(55).Find(&allProducts).Error; err != nil {
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch random products"})
// 		}

// 		// Acak produk
// 		rand.Seed(time.Now().UnixNano())
// 		rand.Shuffle(len(allProducts), func(i, j int) {
// 			allProducts[i], allProducts[j] = allProducts[j], allProducts[i]
// 		})

// 		return c.JSON(fiber.Map{
// 			"recommended_products": allProducts,
// 		})
// 	}

// 	// Cari produk dengan kategori yang sesuai dari histori user
// 	var recommendedProducts []models.Product
// 	if err := database.DB.Where("category IN (?)", categories).Limit(55).Find(&recommendedProducts).Error; err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch recommended products"})
// 	}

// 	// Jika produk kurang dari 55, tambahkan produk acak untuk melengkapi
// 	if len(recommendedProducts) < 55 {
// 		var additionalProducts []models.Product
// 		neededProducts := 55 - len(recommendedProducts)
// 		if err := database.DB.Limit(neededProducts).Find(&additionalProducts).Error; err != nil {
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch additional products"})
// 		}
// 		recommendedProducts = append(recommendedProducts, additionalProducts...)
// 	}

// 	// Acak urutan produk yang direkomendasikan
// 	rand.Seed(time.Now().UnixNano())
// 	rand.Shuffle(len(recommendedProducts), func(i, j int) {
// 		recommendedProducts[i], recommendedProducts[j] = recommendedProducts[j], recommendedProducts[i]
// 	})

// 	// Return hasil rekomendasi yang sudah diacak
// 	return c.JSON(fiber.Map{
// 		"recommended_products": recommendedProducts,
// 	})
// }

// func TestRekomen(c *fiber.Ctx) error{
// 	userId, ok := c.Locals("user_id").(uint)
// 	if !ok {
// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
// 	}

// 	var histories []models.UserHistory
// 	if err := database.DB.Where("user_id=?" userID).
// 	Order("created_at DESC").Limit(20).Find(&histories).Error; err != nil{
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to fetch user history"})
// 	}


// }