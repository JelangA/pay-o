package controller

import (
	"fmt"
	"net/http"
	"os"
	"pay-o/config"
	"pay-o/models"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {

	//body request
	var body struct {
		Name     string `validate:"required"`
		Email    string `validate:"required"`
		Password string `validate:"required"`
		Phone    int    `validate:"required"`
	}

	//cek request
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})
		return
	}

	validate := validator.New()

	if err := validate.Struct(body); err != nil {
		validationErrors := err.(validator.ValidationErrors)

		// Create a slice to hold individual validation error messages
		var errorMessages []string

		for _, fieldError := range validationErrors {
			// Concatenate field and its corresponding validation error into a string
			errorMessage := fmt.Sprintf("%s %s", fieldError.Field(), fieldError.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		// Combine individual validation error messages into a single string
		combinedErrorMessage := strings.Join(errorMessages, ", ")

		// Send the combined validation error message in the JSON response
		c.JSON(http.StatusBadRequest, gin.H{
			"error": combinedErrorMessage,
		})

		return
	}

	//enkripsi password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})

		return
	}

	//Create the user
	user := models.User{
		Name:     body.Name,
		Email:    body.Email,
		Password: string(hash),
		Phone:    body.Phone,
	}
	result := config.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	//generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create token",
		})
		return
	}

	//response with token
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})

}

func Login(c *gin.Context) {
	//body request
	var body struct {
		Email    string `validate:"required"`
		Password string `validate:"required"`
	}

	//cek request
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})
		return
	}

	validate := validator.New()

	if err := validate.Struct(body); err != nil {
		validationErrors := err.(validator.ValidationErrors)

		// Create a slice to hold individual validation error messages
		var errorMessages []string

		for _, fieldError := range validationErrors {
			// Concatenate field and its corresponding validation error into a string
			errorMessage := fmt.Sprintf("%s %s", fieldError.Field(), fieldError.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		// Combine individual validation error messages into a single string
		combinedErrorMessage := strings.Join(errorMessages, ", ")

		// Send the combined validation error message in the JSON response
		c.JSON(http.StatusBadRequest, gin.H{
			"error": combinedErrorMessage,
		})

		return
	}

	//select user by email
	var user models.User
	config.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email or password",
		})
		return
	}

	//compare with password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email or password",
		})
		return
	}

	//generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create token",
		})
		return
	}

	//response with token
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}
