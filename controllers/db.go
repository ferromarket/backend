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
	db.Migrator().DropTable(&models.Pais{}, &models.Region{}, &models.Ciudad{}, &models.Comuna{}, &models.Ferreteria{}, &models.Dia{}, &models.Hora{}, &models.FerreteriaHorario{})
	db.SetupJoinTable(&models.Ferreteria{}, "Dias", &models.FerreteriaHorario{})
	db.AutoMigrate(&models.Hora{})
	db.AutoMigrate(&models.Ciudad{})
	db.AutoMigrate(&models.Comuna{})
	db.AutoMigrate(&models.Dia{})
	db.AutoMigrate(&models.Ferreteria{})
	db.AutoMigrate(&models.Pais{})
	db.AutoMigrate(&models.Region{})

	DBPopulate(db)
}

func DBPopulate(db *gorm.DB) {
	ferreteria := models.Ferreteria{
		Nombre: "Chris's Hardware Store",
		Direccion: "Canto del Valle 1777",
		Descripcion: "Buy something here mother fucker!",
		Comuna: models.Comuna{
			Nombre: "Hualpen",
			Ciudad: models.Ciudad{
				Nombre: "Concepción",
				Region: models.Region{
					Nombre: "Bío Bío",
					Codigo: "VIII",
					Pais: models.Pais{
						Nombre: "Chile",
						Codigo: "CL",
					},
				},
			},
		},
	}

	dias := []models.Dia{
		{Nombre: "Lunes"},
		{Nombre: "Martes"},
		{Nombre: "Miércoles"},
		{Nombre: "Jueves"},
		{Nombre: "Viernes"},
		{Nombre: "Sabado"},
		{Nombre: "Domingo"},
	}

	horas := []models.Hora{
		{Hora: "08:00"},
		{Hora: "09:00"},
		{Hora: "10:00"},
		{Hora: "11:00"},
		{Hora: "12:00"},
		{Hora: "13:00"},
		{Hora: "14:00"},
		{Hora: "15:00"},
		{Hora: "16:00"},
		{Hora: "17:00"},
		{Hora: "18:00"},
		{Hora: "19:00"},
		{Hora: "20:00"},
		{Hora: "21:00"},
		{Hora: "22:00"},
	}

	horarios := []models.FerreteriaHorario{
		{
			FerreteriaID: 1,
			DiaID: 1,
			AbrirID: 1,
			CerrarID: 15,
		},
	}

	db.Create(&ferreteria)
	db.Create(&dias)
	db.Create(&horas)
	db.Create(&horarios)
}

func DBClose(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get gorm db")
	}
	sqlDB.Close()
}
