package orders

import (
	"errors"
	"fmt"
	"time"

	"github.com/jakewd7/internal/modules/products"
	"gorm.io/gorm"
)

type Service interface {
	Checkout(input CheckoutRequest) (OrderResponse, error)
	GetAllOrders() ([]OrderResponse, error)
	GetOrderByID(id string) (OrderResponse, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func generateOrderNumber() string {
	return fmt.Sprintf("INV-%d", time.Now().UnixNano())
}

func mapOrderToResponse(order Order) OrderResponse {
	var itemResponses []OrderItemResponse
	for _, item := range order.OrderItems {
		itemResponses = append(itemResponses, OrderItemResponse{
			ID:        item.ID,
			ProductID: item.ProductID,
			Product: products.ProductResponse{
				ID:          item.Product.ID,
				CategoryID:  item.Product.CategoryID,
				Name:        item.Product.Name,
				Slug:        item.Product.Slug,
				Description: item.Product.Description,
				Price:       item.Product.Price,
				Stock:       item.Product.Stock,
				ImageURL:    item.Product.ImageURL,
				IsActive:    item.Product.IsActive,
			},
			Price:    item.Price,
			Quantity: item.Quantity,
			Subtotal: item.Subtotal,
		})
	}

	return OrderResponse{
		ID:              order.ID,
		OrderNumber:     order.OrderNumber,
		UserID:          order.UserID,
		CustomerName:    order.CustomerName,
		CustomerPhone:   order.CustomerPhone,
		TotalAmount:     order.TotalAmount,
		Status:          order.Status,
		ShippingAddress: order.ShippingAddress,
		OrderItems:      itemResponses,
		CreatedAt:       order.CreatedAt,
		UpdatedAt:       order.UpdatedAt,
	}
}

func (s *service) Checkout(input CheckoutRequest) (OrderResponse, error) {
	db := s.repository.GetDB()
	var createdOrder Order

	err := db.Transaction(func(tx *gorm.DB) error {
		var totalAmount float64
		var orderItems []OrderItem

		for _, itemReq := range input.Items {
			product, err := s.repository.FindProductByID(itemReq.ProductID)
			if err != nil {
				return errors.New("product not found: " + itemReq.ProductID)
			}

			if product.Stock < itemReq.Quantity {
				return errors.New("insufficient stock for product: " + product.Name)
			}

			product.Stock -= itemReq.Quantity
			if err := s.repository.UpdateProductStock(tx, product); err != nil {
				return err
			}

			subtotal := product.Price * float64(itemReq.Quantity)
			totalAmount += subtotal

			orderItems = append(orderItems, OrderItem{
				ProductID: product.ID,
				Price:     product.Price,
				Quantity:  itemReq.Quantity,
				Subtotal:  subtotal,
			})
		}

		createdOrder = Order{
			OrderNumber:     generateOrderNumber(),
			UserID:          input.UserID,
			CustomerName:    input.CustomerName,
			CustomerPhone:   input.CustomerPhone,
			TotalAmount:     totalAmount,
			Status:          "PENDING",
			ShippingAddress: input.ShippingAddress,
			OrderItems:      orderItems,
		}

		if err := s.repository.CreateOrder(tx, &createdOrder); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return OrderResponse{}, err
	}

	fetchedOrder, err := s.repository.FindByID(createdOrder.ID)
	if err != nil {
		return OrderResponse{}, err
	}

	return mapOrderToResponse(*fetchedOrder), nil
}

func (s *service) GetAllOrders() ([]OrderResponse, error) {
	orders, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	var responses []OrderResponse
	for _, o := range orders {
		responses = append(responses, mapOrderToResponse(o))
	}
	return responses, nil
}

func (s *service) GetOrderByID(id string) (OrderResponse, error) {
	order, err := s.repository.FindByID(id)
	if err != nil {
		return OrderResponse{}, err
	}
	return mapOrderToResponse(*order), nil
}
