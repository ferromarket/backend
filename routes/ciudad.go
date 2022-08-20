package routes

import (
	"github.com/ferromarket/backend/controllers"
	"github.com/julienschmidt/httprouter"
)

func CiudadRoutes(router *httprouter.Router) {
	router.GET("/ciudades", controllers.ListCiudades)
	router.GET("/ciudades/:region", controllers.ListCiudades)
	router.GET("/ciudad/:id", controllers.GetCiudad)
}
