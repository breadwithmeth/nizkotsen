package handlers

import (
    "github.com/breadwithmeth/nizkotsen/internal/service"
    "github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
    service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
    return &ProductHandler{service: service}
}

func (h *ProductHandler) GetProducts(c *fiber.Ctx) error {
    products, err := h.service.GetAllProducts()
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(products)
}

// Другие обработчики: GetProduct, CreateProduct, UpdateProduct, DeleteProduct