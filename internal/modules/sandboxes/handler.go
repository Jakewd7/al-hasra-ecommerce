package sandboxes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

// GetPaymentByVA godoc
// @Summary Inquiry Virtual Account status
// @Description Get Virtual Account payment details by VA Number
// @Tags Sandbox
// @Produce json
// @Param va_number path string true "VA Number"
// @Success 200 {object} SandboxPaymentResponse
// @Failure 404 {object} map[string]string
// @Router /sandbox/va/{va_number} [get]
func (h *Handler) GetPaymentByVA(c *gin.Context) {
	vaNumber := c.Param("va_number")
	if vaNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nomor VA wajib diisi"})
		return
	}

	response, err := h.service.GetPaymentByVA(vaNumber)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// PaySandbox godoc
// @Summary Simulate VA Payment (Sandbox)
// @Description Process payment by VA Number and update status from PENDING to PAID
// @Tags Sandbox
// @Accept json
// @Produce json
// @Param request body PaySandboxRequest true "Sandbox Payment Request"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /sandbox/pay [post]
func (h *Handler) PaySandbox(c *gin.Context) {
	var input PaySandboxRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	paymentData, err := h.service.ProcessPayment(input.VANumber)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Pembayaran berhasil",
		"data":    paymentData,
	})
}
