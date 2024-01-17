package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Validate(c *gin.Context) {

	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}