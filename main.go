package main

import (
	"github.com/ferromarket/backend/routes"
)

func main() {
    router := routes.Initialize()
    routes.Serve(router)
}
