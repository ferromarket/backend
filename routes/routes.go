package routes

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/julienschmidt/httprouter"
)

type Message struct {
    Message string
}

func Initialize() (*httprouter.Router) {
    router := httprouter.New()
    router.GET("/", index)
    FerreteriaRoutes(router)
    return router
}

func Serve(router *httprouter.Router) {
    newRouter := handlers.CombinedLoggingHandler(os.Stdout, router)
    newRouter = handlers.CompressHandler(newRouter)
    log.Fatal(http.ListenAndServe(":3001", newRouter))
}

func index(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
    fmt.Fprintf(writer, "This is the FerroMarket API server!")
}
