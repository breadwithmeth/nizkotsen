package routes

import (
    "github.com/breadwithmeth/nizkotsen/internal/api/handlers"
    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, productHandler *handlers.ProductHandler) {
    api := app.Group("/api")
    
    // Маршруты продуктов
    products := api.Group("/products")
    products.Get("/", productHandler.GetProducts)
    // Другие маршруты
}