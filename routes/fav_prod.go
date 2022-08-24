package routes

import (
	"github.com/ferromarket/backend/controllers"
	"github.com/ferromarket/backend/middlewares"
	"github.com/julienschmidt/httprouter"
)

func FavProdRoutes(router *httprouter.Router) {
	router.POST("/favProd", middlewares.Authenticate(controllers.PostFavProd))
	router.GET("/favProd", middlewares.Authenticate(controllers.ListFavProd))
	router.GET("/favProd/:id", middlewares.Authenticate(controllers.GetFavProd))
	router.PUT("/favProd/:id", middlewares.Authenticate(controllers.PutFavProd))
	router.PATCH("/favProd/:id", middlewares.Authenticate(controllers.PatchFavProd))
	router.DELETE("/favProd/:id", middlewares.Authenticate(controllers.DeleteFavProd))
}
