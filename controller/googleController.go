package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

var (
	googleOauthConfig = &oauth2.Config{
		ClientID:     "YOUR_GOOGLE_CLIENT_ID",
		ClientSecret: "YOUR_GOOGLE_CLIENT_SECRET",
		RedirectURL:  "http://your-backend-url/auth/google/callback",
		Scopes:       []string{"openid", "profile", "email"},
		Endpoint:     oauth2.Endpoint{},
	}

	oauthStateString = "random" // state string untuk melindungi dari CSRF
	jwtSecret        = "your_jwt_secret"
)

func testing() {
	router := gin.Default()

	router.GET("/auth/google", handleGoogleLogin)
	router.GET("/auth/google/callback", handleGoogleCallback)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server listening on :%s...\n", port)
	log.Fatal(router.Run(":" + port))
}

func handleGoogleLogin(c *gin.Context) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func handleGoogleCallback(c *gin.Context) {
	state := c.Query("state")
	if state != oauthStateString {
		fmt.Printf("Invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	code := c.Query("code")
	token, err := googleOauthConfig.Exchange(c, code)
	if err != nil {
		fmt.Printf("Code exchange failed with error: %s\n", err)
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	client := googleOauthConfig.Client(c, token)
	userInfo, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		fmt.Printf("Failed to get user info: %s\n", err)
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	defer userInfo.Body.Close()

	var userProfile map[string]interface{}
	if err := json.NewDecoder(userInfo.Body).Decode(&userProfile); err != nil {
		fmt.Printf("Failed to decode user info: %s\n", err)
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	// Proses informasi pengguna dari userProfile sesuai kebutuhan aplikasi
	// ...

	// Contoh: Menghasilkan dan mengirimkan token JWT sebagai respons
	jwtToken := generateJWTToken(userProfile)
	c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("http://your-frontend-url?token=%s", jwtToken))
}

// Fungsi ini hanya contoh untuk menghasilkan token JWT, Anda harus menggantinya dengan implementasi sesuai kebutuhan aplikasi Anda.
func generateJWTToken(userProfile map[string]interface{}) string {
	claims := jwt.MapClaims{
		"sub": userProfile["sub"],
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(), // Contoh: Token berlaku selama 30 hari
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		fmt.Printf("Failed to generate JWT token: %s\n", err)
		return ""
	}

	return tokenString
}