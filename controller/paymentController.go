// midtrans_controller.go

package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

// MidtransController handles Midtrans transactions
type MidtransController struct{}

// NewMidtransController creates a new instance of MidtransController
func NewMidtransController() *MidtransController {
	return &MidtransController{}
}

// CreateTransaction handles the creation of Midtrans transactions
func (mc *MidtransController) CreateTransaction(c *gin.Context) {
	// Parse request data and create a Snap request
	var snapReqData map[string]interface{}
	if err := c.ShouldBindJSON(&snapReqData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Create a Snap request
	snapReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  snapReqData["order_id"].(string),
			GrossAmt: int64(snapReqData["gross_amount"].(float64)),
		},
	}

	// Call the Midtrans CreateTransaction function
	response, err := snap.CreateTransaction(snapReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Midtrans transaction"})
		return
	}

	// Return the Snap response to the client
	c.JSON(http.StatusOK, response)
}
