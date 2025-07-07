package middleware

import (
	"fmt"
	"jwt/initializers"
	"jwt/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	secret := []byte(os.Getenv("SECRET"))
	tokenString, err := c.Cookie("token")
	fmt.Println(tokenString)
	if err != nil {
		fmt.Println("no token")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret,nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["exp"], claims["sub"])
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			fmt.Println("exp")
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		var user models.User
		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			fmt.Println("no user")
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Set("user", user)
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}