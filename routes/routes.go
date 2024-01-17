package routes

import (
	"pay-o/controller"
	"pay-o/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter mengonfigurasi dan mengembalikan instance dari router Gin
func SetupRouter() *gin.Engine {

	r := gin.Default()

	r.ForwardedByClientIP = true
	r.SetTrustedProxies([]string{"127.0.0.1"})

	r.POST("/register", controller.Signup)
	r.POST("/login", controller.Login)
	r.GET("/getuser", middleware.AuthMiddleware, controller.Validate)
	r.GET("/google/login")
	r.GET("/google/callback")

	r.POST("/midtrans/create-transaction", middleware.AuthMiddleware, controller.NewMidtransController().CreateTransaction)

	return r
}
