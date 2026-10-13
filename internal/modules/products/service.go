package products

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type Service interface {
	GetAllProducts() ([]ProductResponse, error)
	GetProductByID(id string) (ProductResponse, error)
	GetProductBySlug(slug string) (ProductResponse, error)
	CreateProduct(input CreateProductRequest) (ProductResponse, error)
	UpdateProduct(id string, input UpdateProductRequest) (ProductResponse, error)
	DeleteProduct(id string) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func generateSlug(title string) string {
	slug := strings.ToLower(title)
	slug = strings.ReplaceAll(slug, " ", "-")
	return slug
}

func mapProductToResponse(product Product) ProductResponse {
	return ProductResponse{
		ID:          product.ID,
		CategoryID:  product.CategoryID,
		Name:        product.Name,
		Slug:        product.Slug,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		ImageURL:    product.ImageURL,
		IsActive:    product.IsActive,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}

func (s *service) GetAllProducts() ([]ProductResponse, error) {
	products, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	responses := make([]ProductResponse, 0)
	for _, p := range products {
		responses = append(responses, mapProductToResponse(p))
	}
	return responses, nil
}

func (s *service) GetProductByID(id string) (ProductResponse, error) {
	product, err := s.repository.FindByID(id)
	if err != nil {
		return ProductResponse{}, err
	}
	return mapProductToResponse(*product), nil
}

func (s *service) GetProductBySlug(slug string) (ProductResponse, error) {
	product, err := s.repository.FindBySlug(slug)
	if err != nil {
		return ProductResponse{}, err
	}
	return mapProductToResponse(*product), nil
}

func (s *service) CreateProduct(input CreateProductRequest) (ProductResponse, error) {
	slug := generateSlug(input.Name)

	existing, _ := s.repository.FindBySlug(slug)
	if existing != nil && existing.ID != "" {
		slug = fmt.Sprintf("%s-%d", slug, time.Now().Unix())
	}

	product := Product{
		CategoryID:  input.CategoryID,
		Name:        input.Name,
		Slug:        slug,
		Description: input.Description,
		Price:       input.Price,
		Stock:       input.Stock,
		ImageURL:    input.ImageURL,
		IsActive:    true,
	}

	if err := s.repository.Create(&product); err != nil {
		return ProductResponse{}, err
	}

	return mapProductToResponse(product), nil
}

func (s *service) UpdateProduct(id string, input UpdateProductRequest) (ProductResponse, error) {
	product, err := s.repository.FindByID(id)
	if err != nil {
		return ProductResponse{}, errors.New("product not found")
	}

	if input.CategoryID != "" {
		product.CategoryID = input.CategoryID
	}
	if input.Name != "" {
		product.Name = input.Name
		newSlug := generateSlug(input.Name)
		if newSlug != product.Slug {
			existing, _ := s.repository.FindBySlug(newSlug)
			if existing != nil && existing.ID != product.ID {
				newSlug = fmt.Sprintf("%s-%d", newSlug, time.Now().Unix())
			}
			product.Slug = newSlug
		}
	}
	if input.Description != "" {
		product.Description = input.Description
	}
	if input.Price != nil {
		product.Price = *input.Price
	}
	if input.Stock != nil {
		product.Stock = *input.Stock
	}
	if input.ImageURL != "" {
		product.ImageURL = input.ImageURL
	}
	if input.IsActive != nil {
		product.IsActive = *input.IsActive
	}

	product.UpdatedAt = time.Now()

	if err := s.repository.Update(product); err != nil {
		return ProductResponse{}, err
	}

	return mapProductToResponse(*product), nil
}

func (s *service) DeleteProduct(id string) error {
	return s.repository.Delete(id)
}
