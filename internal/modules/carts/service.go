package carts

import (
	"github.com/jakewd7/internal/modules/products"
)

type Service interface {
	AddToCart(input AddToCartRequest) (CartResponse, error)
	GetAllCarts() ([]CartResponse, error)
	DeleteCart(id string) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) AddToCart(input AddToCartRequest) (CartResponse, error) {
	cart := Cart{
		UserID:    input.UserID,
		ProductID: input.ProductID,
		Quantity:  input.Quantity,
	}

	if err := s.repository.Create(&cart); err != nil {
		return CartResponse{}, err
	}

	fetchedCart, err := s.repository.FindByID(cart.ID)
	if err != nil {
		return CartResponse{}, err
	}

	return CartResponse{
		ID:        fetchedCart.ID,
		UserID:    fetchedCart.UserID,
		ProductID: fetchedCart.ProductID,
		Product: products.ProductResponse{
			ID:          fetchedCart.Product.ID,
			CategoryID:  fetchedCart.Product.CategoryID,
			Name:        fetchedCart.Product.Name,
			Slug:        fetchedCart.Product.Slug,
			Description: fetchedCart.Product.Description,
			Price:       fetchedCart.Product.Price,
			Stock:       fetchedCart.Product.Stock,
			ImageURL:    fetchedCart.Product.ImageURL,
			IsActive:    fetchedCart.Product.IsActive,
		},
		Quantity:  fetchedCart.Quantity,
		CreatedAt: fetchedCart.CreatedAt,
		UpdatedAt: fetchedCart.UpdatedAt,
	}, nil
}

func (s *service) GetAllCarts() ([]CartResponse, error) {
	carts, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	var responses []CartResponse
	for _, c := range carts {
		responses = append(responses, CartResponse{
			ID:        c.ID,
			UserID:    c.UserID,
			ProductID: c.ProductID,
			Product: products.ProductResponse{
				ID:          c.Product.ID,
				CategoryID:  c.Product.CategoryID,
				Name:        c.Product.Name,
				Slug:        c.Product.Slug,
				Description: c.Product.Description,
				Price:       c.Product.Price,
				Stock:       c.Product.Stock,
				ImageURL:    c.Product.ImageURL,
				IsActive:    c.Product.IsActive,
			},
			Quantity:  c.Quantity,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		})
	}
	return responses, nil
}

func (s *service) DeleteCart(id string) error {
	return s.repository.Delete(id)
}
