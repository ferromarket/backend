package routes

import (
	"github.com/ferromarket/backend/controllers"
	"github.com/julienschmidt/httprouter"
)

func RepartidorRoutes(router *httprouter.Router) {
	router.POST("/repartidor", controllers.PostRepartidor)
	router.GET("/repartidores", controllers.ListRepartidores)
	router.GET("/repartidor/:id", controllers.GetRepartidor)
	router.PUT("/repartidor/:id", controllers.PutRepartidor)
	router.PATCH("/repartidor/:id", controllers.PatchRepartidor)
	router.DELETE("/repartidor/:id", controllers.DeleteRepartidor)
}
