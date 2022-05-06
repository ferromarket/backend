package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ferromarket/backend/routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

type Message struct {
    Message string
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprintf(w, "This is the golang api server!")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    datas := Message{}
    datas.Message = "Hello, " + ps.ByName("name")

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(datas)
    //fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
    router := httprouter.New()
    router.GET("/", Index)
    router.GET("/hello/:name", Hello)

    routes.PrintBaz()

    /*dbUser, exists := os.LookupEnv("MYSQL_USER")
    if !exists {
        dbUser = "user"
    }
    dbPass, exists := os.LookupEnv("MYSQL_PASSWORD")
    if !exists {
        dbPass = "pass"
    }
    dbHost, exists := os.LookupEnv("MYSQL_HOST_IP")
    if !exists {
        dbHost = "localhost"
    }
    dbName, exists := os.LookupEnv("MYSQL_DATABASE")
    if !exists {
        dbName = "database"
    }

    db, err := sql.Open("mysql", dbUser + ":" + dbPass + "@tcp(" + dbHost + ")/" + dbName)
    if err != nil {
        panic(err)
    }
    db.SetConnMaxLifetime(time.Minute * 3)
    db.SetMaxOpenConns(10)
    db.SetMaxIdleConns(10)

    defer db.Close()

    insert, err := db.Query("INSERT INTO books_reviews (book_name, book_review) VALUES ( 'TEST1', 'TEST2' )")

    if err != nil {
        panic(err.Error())
    }

    defer insert.Close()

    log.Fatal(http.ListenAndServe(":3001", router))
}
