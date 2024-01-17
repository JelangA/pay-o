package main

import (
	"log"
	"os"
	"pay-o/config"
	"pay-o/routes"

	"github.com/midtrans/midtrans-go"
)



func init(){
	config.LoadEnvVariables()
	config.DBconnection()
	config.SyncDatabase()

	midtrans.ClientKey = os.Getenv("CLIENTKEY")
	midtrans.ServerKey = os.Getenv("SERVERKEY")
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}
	
	
	router := routes.SetupRouter()
	log.Printf("Server listening on http://localhost%s\n", port)
	router.Run(port) // listen and serve on 0.0.0.0:8080
	
}

// package main
// import (
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"strings"
// 	// "github.com/gin-gonic/gin"
// )
// func main() {
// 	url := "https://app.sandbox.midtrans.com/snap/v1/transactions"
// 	// data := 
// 	payload := strings.NewReader("{\"transaction_details\":{\"order_id\":\"order-id\",\"gross_amount\":10000},\"credit_card\":{\"secure\":true}}")
// 	req, _ := http.NewRequest("POST", url, payload)
// 	req.Header.Add("accept", "application/json")
// 	req.Header.Add("content-type", "application/json")
// 	res, _ := http.DefaultClient.Do(req)
// 	defer res.Body.Close()
// 	body, _ := io.ReadAll(res.Body)
// 	fmt.Println(string(body))
// }