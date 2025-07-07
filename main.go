package main

import (
	"jwt/controllers"
	"jwt/initializers"
	"jwt/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvFile()
	initializers.ConnectToDB()
	initializers.Migrate()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)

	r.GET("/private", middleware.RequireAuth, controllers.Private)

	r.Run()
}