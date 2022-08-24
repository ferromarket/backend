package routes

import (
	"github.com/ferromarket/backend/controllers"
	"github.com/ferromarket/backend/middlewares"
	"github.com/julienschmidt/httprouter"
)

func FerreteriaRoutes(router *httprouter.Router) {
	router.POST("/ferreteria", middlewares.Authenticate(controllers.PostFerreteria))
	router.GET("/ferreterias", controllers.ListFerreterias)
	router.GET("/ferreteria/:id", controllers.GetFerreteria)
	router.PUT("/ferreteria/:id", middlewares.Authenticate(controllers.PutFerreteria))
	router.PATCH("/ferreteria/:id", middlewares.Authenticate(controllers.PatchFerreteria))
	router.DELETE("/ferreteria/:id", middlewares.Authenticate(controllers.DeleteFerreteria))
}
