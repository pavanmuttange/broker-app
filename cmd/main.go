package main

import (
	"brokerApp/config"
	"brokerApp/internal/handler"
	"brokerApp/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	db := config.ConnectDB()
	r := gin.Default()

	r.GET("/health", handler.HealthCheck)

	r.POST("/signup", handler.SignUp(db))
	r.POST("/login", handler.Login(db))
	r.POST("/refresh-token", handler.RefreshToken)

	auth := r.Group("/")
	auth.Use(middleware.JWTAuthMiddleware())
	{
		auth.GET("/holdings", handler.GetHoldings)
		auth.GET("/orderbook", handler.GetOrderbook)
		auth.GET("/positions", handler.GetPositions)
	}

	r.Run()
}
