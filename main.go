package main

import (
	"fmt"

	"github.com/ferromarket/backend/controllers"
	"github.com/ferromarket/backend/routes"
)

func main() {
    fmt.Print("Connecting to database ... ")
    db := controllers.DBConnect()
    fmt.Println("DONE")

    fmt.Print("AutoMigrating database ... ")
    controllers.DBAutoMigrate(db)
    fmt.Println("DONE")

    fmt.Print("Closing database ... ")
    controllers.DBClose(db)
    fmt.Println("DONE")

    router := routes.Initialize()
    fmt.Println("Serving routes ... DONE")
    routes.Serve(router)
}
