package routes

import (
	"github.com/ferromarket/backend/controllers"
	"github.com/julienschmidt/httprouter"
)

func CategoriaRoutes(router *httprouter.Router) {
	router.GET("/categorias", controllers.ListCategorias)
	router.GET("/categoria/:id", controllers.GetCategoria)
}
