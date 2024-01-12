// controller/payment_controller.go

package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func initMidtransClient() {}

func ShowPaymentPage(c *gin.Context) {
	initMidtransClient()
	c.HTML(http.StatusOK, "payment_page.html", gin.H{
	})
}

func ChargePayment(c *gin.Context) {
	initMidtransClient()
	// Menyiapkan payload pembayaran untuk Midtrans
	paymentPayload := map[string]interface{}{
		"payment_type": "credit_card",
		"transaction_details": map[string]interface{}{
			"order_id": "ORDER123",
			"gross_amount": 100000,
		},
		"credit_card": map[string]interface{}{
			"token_id": "YOUR_CREDIT_CARD_TOKEN", // Token kartu kredit yang telah di-generate sebelumnya
		},
	}

	// Melakukan request ke API Midtrans
	apiURL := "https://api.sandbox.midtrans.com/v2/charge"
	apiKey := "YOUR_MIDTRANS_SERVER_KEY"
	response, err := callMidtransAPI(apiURL, apiKey, paymentPayload)

	if response != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to charge payment: %s", err.Error()),
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to charge payment: %s", err.Error()),
		})
		return
	}

	// Melakukan sesuatu dengan response, misalnya menyimpan informasi pembayaran ke database
	// ...

	c.JSON(http.StatusOK, gin.H{
		"message": "Payment successful!",
		// Kirim data atau pesan lain ke frontend
		// ...
	})
}

func callMidtransAPI(apiURL, apiKey string, payload map[string]interface{}) (*http.Response, error) {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+apiKey)

	client := &http.Client{}
	return client.Do(req)
}
