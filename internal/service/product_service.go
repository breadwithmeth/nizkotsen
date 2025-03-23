package service

import (
    "github.com/breadwithmeth/nizkotsen/internal/models"
    "github.com/breadwithmeth/nizkotsen/internal/repository"
)

type ProductService struct {
    repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
    return &ProductService{repo: repo}
}

func (s *ProductService) GetAllProducts() ([]models.Product, error) {
    return s.repo.GetAll()
}

// Другие бизнес-методы