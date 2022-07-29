package routes

import (
	"github.com/ferromarket/backend/controllers"
	"github.com/julienschmidt/httprouter"
)

func UsuarioRoutes (router *httprouter.Router){
	router.POST("/usuario", controllers.PostUsuario)
	router.GET("/usuarios", controllers.ListUsuarios)
	router.GET("/usuario/:id", controllers.GetUsuario)
	router.PUT("/usuario/:id", controllers.PutUsuario)
	router.PATCH("/usuario/:id", controllers.PatchUsuario)
	router.DELETE("/usuario/:id", controllers.DeleteUsuario)
}
