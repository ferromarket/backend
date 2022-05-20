package main

import (
	"fmt"
	"os"

	"github.com/ferromarket/backend/database"
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
    gdb := database.Connect()
    fmt.Println("DONE")

    if (dropDB) {
        fmt.Print("Dropping database ... ")
        database.DropAll(gdb)
        fmt.Println("DONE")
    }

    fmt.Print("AutoMigrating database ... ")
    database.AutoMigrate(gdb)
    fmt.Println("DONE")

    if (populateDB) {
        fmt.Print("Populating database ... ")
        database.Populate(gdb)
        fmt.Println("DONE")
    }

    fmt.Print("Closing database ... ")
    database.Close(gdb)
    fmt.Println("DONE")

    router := routes.Initialize()
    fmt.Println("Serving routes ... DONE")
    routes.Serve(router)
}
