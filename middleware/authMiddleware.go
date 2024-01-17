package middleware

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "pay-o/config"
    "pay-o/models"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware(c *gin.Context) {
    // get token from Authorization header
    tokenString := c.GetHeader("Authorization")

    if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
    // validation
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }

        return []byte(os.Getenv("SECRET")), nil
    })
    if err != nil {
        log.Println("Error parsing token:", err)
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
        return
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        if expirationTime, ok := claims["exp"].(float64); ok {
            if time.Now().Unix() > int64(expirationTime) {
                c.JSON(http.StatusUnauthorized, gin.H{"error": "Token has expired"})
				c.Abort()
                return
            }
        } else {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token expiration"})
			c.Abort()
            return
        }

        var user models.User
        config.DB.First(&user, claims["sub"])

        if user.ID == 0 {
            c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			c.Abort()
            return
        }

        c.Set("user", user)
        c.Next()
    } else {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		c.Abort()
        return
    }

    fmt.Println("Generated Token:", tokenString)
}