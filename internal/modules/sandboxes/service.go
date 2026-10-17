package sandboxes

import (
	"errors"
	"time"

	"github.com/jakewd7/internal/modules/orders"
	"gorm.io/gorm"
)

type Service interface {
	GetPaymentByVA(vaNumber string) (SandboxPaymentResponse, error)
	ProcessPayment(vaNumber string) (SandboxPaymentResponse, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) GetPaymentByVA(vaNumber string) (SandboxPaymentResponse, error) {
	payment, err := s.repository.FindByVANumber(vaNumber)
	if err != nil {
		return SandboxPaymentResponse{}, errors.New("nomor VA tidak ditemukan")
	}

	order, err := s.repository.GetOrderWithDetails(payment.OrderID)
	if err != nil {
		return SandboxPaymentResponse{}, errors.New("data order tidak ditemukan")
	}

	expiresIn := int64(time.Until(payment.ExpiredAt).Seconds())
	if expiresIn < 0 {
		expiresIn = 0
	}

	var itemResponses []orders.OrderItemResponse
	for _, item := range order.OrderItems {
		itemResponses = append(itemResponses, orders.OrderItemResponse{
			ProductID:   item.ProductID,
			ProductName: item.Product.Name,
			ImageURL:    item.Product.ImageURL,
			Price:       item.Price,
			Quantity:    item.Quantity,
			Subtotal:    item.Subtotal,
		})
	}

	return SandboxPaymentResponse{
		PaymentID:     payment.ID,
		OrderID:       payment.OrderID,
		OrderNumber:   order.OrderNumber,
		PaymentMethod: payment.PaymentMethod,
		VANumber:      payment.VANumber,
		Amount:        payment.Amount,
		Status:        payment.Status,
		ExpiredAt:     payment.ExpiredAt,
		ExpiresIn:     expiresIn,
		PaidAt:        payment.PaidAt,
		Items:         itemResponses,
	}, nil
}

func (s *service) ProcessPayment(vaNumber string) (SandboxPaymentResponse, error) {
	db := s.repository.GetDB()
	var updatedPayment orders.Payment
	var associatedOrder orders.Order

	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Set("gorm:query_option", "FOR UPDATE").
			Where("va_number = ?", vaNumber).
			First(&updatedPayment).Error; err != nil {
			return errors.New("nomor VA tidak ditemukan")
		}

		if updatedPayment.Status == "PAID" {
			return errors.New("pembayaran sudah lunas sebelumnya")
		}

		if time.Now().After(updatedPayment.ExpiredAt) {
			updatedPayment.Status = "EXPIRED"
			tx.Save(&updatedPayment)
			return errors.New("masa berlaku nomor VA telah kedaluwarsa")
		}

		if err := tx.Set("gorm:query_option", "FOR UPDATE").
			Where("id = ?", updatedPayment.OrderID).
			First(&associatedOrder).Error; err != nil {
			return errors.New("data order tidak ditemukan")
		}

		now := time.Now()
		updatedPayment.Status = "PAID"
		updatedPayment.PaidAt = &now
		if err := tx.Save(&updatedPayment).Error; err != nil {
			return err
		}

		associatedOrder.Status = "PAID"
		associatedOrder.UpdatedAt = now
		if err := tx.Save(&associatedOrder).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return SandboxPaymentResponse{}, err
	}

	return s.GetPaymentByVA(vaNumber)
}
