package service

import (
	"backend/internal/dto"
	"backend/internal/model"
	"context"
	"time"
)

type ProductService interface {
	CreateProduct(ctx context.Context, req *dto.CreateProductRequest) (*dto.ProductResponse, error)
	GetProduct(ctx context.Context, id int) (*dto.ProductResponse, error)
	ListProducts(ctx context.Context) ([]dto.ProductResponse, error)
}

type productService struct {
	repo  repository.ProductRepository
	cache repository.Cache
}

func NewProductService(repo repository.ProductRepository, cache repository.Cache) ProductService {
	return &productService{repo: repo, cache: cache}
}

func (s *productService) CreateProduct(ctx context.Context, req *dto.CreateProductRequest) (*dto.ProductResponse, error) {
	product := model.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.repo.Create(ctx, &product); err != nil {
		return nil, err
	}

	// Invalidate cache
	s.cache.Delete("products")

	return &dto.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
	}, nil
}

// Similar implementations for GetProduct and ListProducts
