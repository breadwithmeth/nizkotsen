package main

import (
    "log"

    "github.com/breadwithmeth/nizkotsen/internal/api/handlers"
    "github.com/breadwithmeth/nizkotsen/internal/api/routes"
    "github.com/breadwithmeth/nizkotsen/internal/database"
    "github.com/breadwithmeth/nizkotsen/internal/repository"
    "github.com/breadwithmeth/nizkotsen/internal/service"
    "github.com/gofiber/fiber/v2"
)

func main() {
    // Инициализация базы данных
    db := database.Connect()
    defer db.Close()

    // Инициализация репозиториев
    productRepo := repository.NewProductRepository(db)

    // Инициализация сервисов
    productService := service.NewProductService(productRepo)

    // Инициализация обработчиков
    productHandler := handlers.NewProductHandler(productService)

    // Настройка сервера
    app := fiber.New()

    // Настройка маршрутов
    routes.SetupRoutes(app, productHandler)

    // Запуск сервера
    log.Fatal(app.Listen(":3000"))
}