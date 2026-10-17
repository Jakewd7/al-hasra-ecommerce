package orders

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

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

func generateBCAVANumber() string {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("88001%08d", rng.Intn(100000000))
}

func mapOrderToResponse(order Order) OrderResponse {
	var itemResponses []OrderItemResponse
	for _, item := range order.OrderItems {
		itemResponses = append(itemResponses, OrderItemResponse{
			ProductID:   item.ProductID,
			ProductName: item.Product.Name,
			ImageURL:    item.Product.ImageURL,
			Price:       item.Price,
			Quantity:    item.Quantity,
			Subtotal:    item.Subtotal,
		})
	}

	var paymentResp *PaymentResponse
	if order.Payment != nil {
		expiresIn := int64(time.Until(order.Payment.ExpiredAt).Seconds())
		if expiresIn < 0 {
			expiresIn = 0
		}
		paymentResp = &PaymentResponse{
			ID:            order.Payment.ID,
			PaymentMethod: order.Payment.PaymentMethod,
			VANumber:      order.Payment.VANumber,
			Amount:        order.Payment.Amount,
			Status:        order.Payment.Status,
			ExpiredAt:     order.Payment.ExpiredAt,
			ExpiresIn:     expiresIn,
			PaidAt:        order.Payment.PaidAt,
		}
	}

	return OrderResponse{
		ID:          order.ID,
		OrderNumber: order.OrderNumber,
		TotalAmount: order.TotalAmount,
		Status:      order.Status,
		OrderItems:  itemResponses,
		Payment:     paymentResp,
		CreatedAt:   order.CreatedAt,
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
			OrderNumber: generateOrderNumber(),
			TotalAmount: totalAmount,
			Status:      "PENDING",
			OrderItems:  orderItems,
		}

		if err := s.repository.CreateOrder(tx, &createdOrder); err != nil {
			return err
		}

		expiredAt := time.Now().Add(24 * time.Hour)
		payment := Payment{
			OrderID:       createdOrder.ID,
			PaymentMethod: "BCA_VA",
			VANumber:      generateBCAVANumber(),
			Amount:        totalAmount,
			Status:        "PENDING",
			ExpiredAt:     expiredAt,
		}

		if err := s.repository.CreatePayment(tx, &payment); err != nil {
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
