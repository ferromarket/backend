package main

import (
	"fmt"
	"os"

	"github.com/ferromarket/backend/controllers"
	"github.com/ferromarket/backend/routes"
)

func main() {
    allArgs := os.Args[1:]

    dropDB := false
    populateDB := false

    for _, element := range allArgs {
        if (element == "drop") {
            dropDB = true
        }

        if (element == "populate") {
            populateDB = true
        }
    }

    fmt.Print("Connecting to database ... ")
    db := controllers.DBConnect()
    fmt.Println("DONE")

    if (dropDB) {
        fmt.Print("Dropping database ... ")
        controllers.DBDropAll(db)
        fmt.Println("DONE")
    }

    fmt.Print("AutoMigrating database ... ")
    controllers.DBAutoMigrate(db)
    fmt.Println("DONE")

    if (populateDB) {
        fmt.Print("Populating database ... ")
        controllers.DBPopulate(db)
        fmt.Println("DONE")
    }

    fmt.Print("Closing database ... ")
    controllers.DBClose(db)
    fmt.Println("DONE")

    router := routes.Initialize()
    fmt.Println("Serving routes ... DONE")
    routes.Serve(router)
}
