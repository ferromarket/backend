package routes

import (
	"github.com/ferromarket/backend/controllers"
	"github.com/julienschmidt/httprouter"
)

func ProductoRoutes(router *httprouter.Router) {
	router.POST("/producto", controllers.PostProducto)
	router.GET("/productos", controllers.ListProductos)
	router.GET("/producto/:id", controllers.GetProducto)
	router.PUT("/producto/:id", controllers.PutProducto)
	router.PATCH("/producto/:id", controllers.PatchProducto)
	router.DELETE("/producto/:id", controllers.DeleteProducto)
}
