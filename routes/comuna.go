package routes

import (
	"github.com/ferromarket/backend/controllers"
	"github.com/julienschmidt/httprouter"
)

func ComunaRoutes(router *httprouter.Router) {
	router.GET("/comunas", controllers.ListComunas)
	router.GET("/comunas/:ciudad", controllers.ListComunas)
	router.GET("/comuna/:id", controllers.GetComuna)
}
