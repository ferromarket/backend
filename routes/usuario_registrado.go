package routes

import (
	"github.com/ferromarket/backend/controllers"
	"github.com/ferromarket/backend/middlewares"
	"github.com/julienschmidt/httprouter"
)

func UsuarioRoutes(router *httprouter.Router) {
	router.POST("/usuario", controllers.PostUsuario)
	router.GET("/usuarios", controllers.ListUsuarios)
	router.GET("/usuario/:id", controllers.GetUsuario)
	router.PUT("/usuario/:id", middlewares.Authenticate(controllers.PutUsuario))
	router.PATCH("/usuario/:id", middlewares.Authenticate(controllers.PatchUsuario))
	router.DELETE("/usuario/:id", middlewares.Authenticate(controllers.DeleteUsuario))
}
