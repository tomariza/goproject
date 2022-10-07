package main

import (

	"tutorial/handler"
	"tutorial/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	
	r := gin.Default()
	r.POST("/login", handler.LoginHandler)

	r.GET("/api/product",
		middleware.ValidateToken(),
		middleware.Authorization([]int{4}),
		handler.GetAllProducts)
	

	r.Run("localhost:8080")
}