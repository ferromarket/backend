package routes

import (
	"github.com/ferromarket/backend/controllers"
	"github.com/julienschmidt/httprouter"
)

func HoraRoutes(router *httprouter.Router) {
	router.GET("/horas", controllers.ListHoras)
}
