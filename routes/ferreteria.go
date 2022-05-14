package routes

import (
	"github.com/ferromarket/backend/controllers"
	"github.com/julienschmidt/httprouter"
)

func FerreteriaRoutes(router *httprouter.Router) {
	router.POST("/ferreteria", controllers.PostFerreteria)
	router.GET("/ferreterias", controllers.ListFerreterias)
	router.GET("/ferreteria/:id", controllers.GetFerreteria)
	router.PUT("/ferreteria/:id", controllers.PutFerreteria)
	router.PATCH("/ferreteria/:id", controllers.PatchFerreteria)
	router.DELETE("/ferreteria/:id", controllers.DeleteFerreteria)
}
