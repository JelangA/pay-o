// midtrans_controller.go

package controller

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

// MidtransController handles Midtrans transactions
type MidtransController struct {
	orderIDCounter int
	mu             sync.Mutex
}

// NewMidtransController creates a new instance of MidtransController
func NewMidtransController() *MidtransController {
	return &MidtransController{}
}

// CreateTransaction handles the creation of Midtrans transactions
func (mc *MidtransController) CreateTransaction(c *gin.Context) {
	// Lock to ensure that orderIDCounter is accessed and updated atomically
	mc.mu.Lock()
	defer mc.mu.Unlock()

	// Increment the order ID counter
	mc.orderIDCounter++

	// Parse request data and create a Snap request
	var reqData map[string]interface{}
	if err := c.ShouldBindJSON(&reqData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Create a Snap request
	ChargeReq := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeGopay,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  generateOrderID(mc.orderIDCounter),
			GrossAmt: int64(reqData["gross_amount"].(float64)),
		},
	}

	// Call the Midtrans CreateTransaction function
	response, err := coreapi.ChargeTransaction(ChargeReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Midtrans transaction"})
		return
	}

	// Return the Snap response to the client
	c.JSON(http.StatusOK, response)
}

// generateOrderID generates a unique order ID based on the counter
func generateOrderID(counter int) string {
	return fmt.Sprintf("ORDER%d", counter)
}
