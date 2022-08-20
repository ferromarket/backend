package routes

import (
	"github.com/ferromarket/backend/controllers"
	"github.com/julienschmidt/httprouter"
)

func VehiculoRoutes(router *httprouter.Router) {
	router.POST("/vehiculo", controllers.PostVehiculo)
	router.GET("/vehiculos", controllers.ListVehiculos)
	router.GET("/vehiculo/:id", controllers.GetVehiculo)
	router.PUT("/vehiculo/:id", controllers.PutVehiculo)
	router.PATCH("/vehiculo/:id", controllers.PatchVehiculo)
	router.DELETE("/vehiculo/:id", controllers.DeleteVehiculo)
}
