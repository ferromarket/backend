package controllers

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func dbConnect() (*gorm.DB) {
	dbUser, exists := os.LookupEnv("MYSQL_USER")
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

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dbUser + ":" + dbPass + "@tcp(" + dbHost + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultStringSize: 256,
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func dbClose(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get gorm db")
	}
	sqlDB.Close()
}
