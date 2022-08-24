package routes

import (
	"github.com/ferromarket/backend/controllers"
	"github.com/ferromarket/backend/middlewares"
	"github.com/julienschmidt/httprouter"
)

func FavProdRoutes(router *httprouter.Router) {
	router.POST("/favprod", middlewares.Authenticate(controllers.PostFavProd))
	router.GET("/favprod", middlewares.Authenticate(controllers.ListFavProd))
	router.GET("/favprod/:id", middlewares.Authenticate(controllers.GetFavProd))
	router.PUT("/favprod/:id", middlewares.Authenticate(controllers.PutFavProd))
	router.PATCH("/favprod/:id", middlewares.Authenticate(controllers.PatchFavProd))
	router.DELETE("/favprod/:id", middlewares.Authenticate(controllers.DeleteFavProd))
}
