package sandboxes

import (
	"time"

	"github.com/jakewd7/internal/modules/orders"
)

type PaySandboxRequest struct {
	VANumber string `json:"va_number" binding:"required" example:"8800112345678"`
}

type SandboxPaymentResponse struct {
	PaymentID     string                     `json:"payment_id"`
	OrderID       string                     `json:"order_id"`
	OrderNumber   string                     `json:"order_number"`
	PaymentMethod string                     `json:"payment_method"`
	VANumber      string                     `json:"va_number"`
	Amount        float64                    `json:"amount"`
	Status        string                     `json:"status"`
	ExpiredAt     time.Time                  `json:"expired_at"`
	ExpiresIn     int64                      `json:"expires_in"`
	PaidAt        *time.Time                 `json:"paid_at,omitempty"`
	Items         []orders.OrderItemResponse `json:"items"`
}
