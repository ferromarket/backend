package controllers

import (
	"os"

	"github.com/ferromarket/backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func DBConnect() (*gorm.DB) {
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
        dbName = "ferromarket"
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dbUser + ":" + dbPass + "@tcp(" + dbHost + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultStringSize: 256,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func DBAutoMigrate(db *gorm.DB) {
	db.SetupJoinTable(&models.Ferreteria{}, "Horarios", &models.FerreteriaHorario{})
	db.AutoMigrate(&models.Ciudad{})
	db.AutoMigrate(&models.Comuna{})
	db.AutoMigrate(&models.Dia{})
	db.AutoMigrate(&models.Ferreteria{})
	db.AutoMigrate(&models.Hora{})
	db.AutoMigrate(&models.FerreteriaHorario{})
	db.AutoMigrate(&models.Pais{})
	db.AutoMigrate(&models.Region{})
}

func DBClose(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get gorm db")
	}
	sqlDB.Close()
}
