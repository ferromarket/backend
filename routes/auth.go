package routes

import (
	"github.com/ferromarket/backend/controllers"
	"github.com/ferromarket/backend/middlewares"
	"github.com/julienschmidt/httprouter"
)

func AuthRoutes(router *httprouter.Router) {
	router.POST("/login", controllers.Login)
	router.POST("/auth", middlewares.Authenticate(controllers.AuthenticateUser))
}
