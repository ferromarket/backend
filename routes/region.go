package routes

import (
	"github.com/ferromarket/backend/controllers"
	"github.com/julienschmidt/httprouter"
)

func RegionRoutes(router *httprouter.Router) {
	router.GET("/regiones", controllers.ListRegiones)
	router.GET("/regiones/:pais", controllers.ListRegiones)
	router.GET("/region/:id", controllers.GetRegion)
}
