package main

import (
	"rest_api_go/controllers/productcontroller"
	"rest_api_go/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.Conn()

	r.GET("/api/products", productcontroller.Index)
	r.GET("/api/product/:id", productcontroller.Show)
	r.POST("/api/product", productcontroller.Create)
	r.PUT("/api/product/:id", productcontroller.Update)
	r.DELETE("/api/products", productcontroller.Delete)

	r.Run()

}
