package main

import (
	handlers "charge_service/api"
	"charge_service/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middleware.CorsMiddleware)

	r.POST("/api/charges", handlers.CreateTransaction)
	r.GET("/api/charges", handlers.GetTransactions)
	r.PUT("/api/charges/:id", handlers.PutTransaction)

	r.Run(":7070")
}

