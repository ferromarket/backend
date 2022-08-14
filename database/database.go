package database

import (
	"os"

	"github.com/ferromarket/backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Connect() (*gorm.DB) {
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

	gdb, err := gorm.Open(mysql.New(mysql.Config{
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

	return gdb
}

func AutoMigrate(gdb *gorm.DB) {
	gdb.AutoMigrate(&models.Hora{})
	gdb.AutoMigrate(&models.FerreteriaHorario{})
	gdb.AutoMigrate(&models.Ciudad{})
	gdb.AutoMigrate(&models.Comuna{})
	gdb.AutoMigrate(&models.Dia{})
	gdb.AutoMigrate(&models.Ferreteria{})
	gdb.AutoMigrate(&models.Pais{})
	gdb.AutoMigrate(&models.Region{})
	gdb.AutoMigrate(&models.Vehiculo{})
	gdb.AutoMigrate(&models.Repartidor{})
	gdb.AutoMigrate(&models.Usuario{})
}

func DropAll(gdb *gorm.DB) {
	gdb.Migrator().DropTable(
		&models.Pais{},
		&models.Region{},
		&models.Ciudad{},
		&models.Comuna{},
		&models.Ferreteria{},
		&models.Dia{},
		&models.Hora{},
		&models.FerreteriaHorario{},
		&models.Usuario{},
		&models.Repartidor{},
		&models.Usuario{})
}

func Populate(gdb *gorm.DB) {
	populateFerreteria(gdb)
}

func populateFerreteria(gdb *gorm.DB) {
	pais := models.Pais{
		ID: 1,
		Nombre: "Chile",
		Codigo: "CL",
	}

	regiones := []models.Region{
		/*{
			ID: 1,
			Nombre: "Tarapacá",
			Codigo: "I",
			PaisID: 1,
		},
		{
			ID: 2,
			Nombre: "Antofagasta",
			Codigo: "II",
			PaisID: 1,
		},
		{
			ID: 3,
			Nombre: "Atacama",
			Codigo: "III",
			PaisID: 1,
		},
		{
			ID: 4,
			Nombre: "Coquimbo",
			Codigo: "IV",
			PaisID: 1,
		},
		{
			ID: 5,
			Nombre: "Valparaíso",
			Codigo: "V",
			PaisID: 1,
		},
		{
			ID: 6,
			Nombre: "Libertador Gral. B.",
			Codigo: "VI",
			PaisID: 1,
		},
		{
			ID: 7,
			Nombre: "Maule",
			Codigo: "VII",
			PaisID: 1,
		},*/
		{
			ID: 8,
			Nombre: "Bíobío",
			Codigo: "VIII",
			PaisID: 1,
		},
		/*{
			ID: 9,
			Nombre: "La Araucanía",
			Codigo: "IX",
			PaisID: 1,
		},
		{
			ID: 10,
			Nombre: "Los Lagos",
			Codigo: "X",
			PaisID: 1,
		},
		{
			ID: 11,
			Nombre: "Aysén del General Carlos",
			Codigo: "XI",
			PaisID: 1,
		},
		{
			ID: 12,
			Nombre: "Magallanes y de la Antártica",
			Codigo: "XII",
			PaisID: 1,
		},
		{
			ID: 13,
			Nombre: "Metropolitana de Santiago",
			Codigo: "RM",
			PaisID: 1,
		},
		{
			ID: 14,
			Nombre: "Los Ríos",
			Codigo: "XIV",
			PaisID: 1,
		},
		{
			ID: 15,
			Nombre: "Arica Parinacota",
			Codigo: "XV",
			PaisID: 1,
		},*/
	}

	ciudades := []models.Ciudad{
		{
			ID: 1,
			Nombre: "Concepción",
			RegionID: 8,
		},
		{
			ID: 2,
			Nombre: "Chillán",
			RegionID: 8,
		},
		{
			ID: 3,
			Nombre: "Talcahuano",
			RegionID: 8,
		},
	}

	comunas := []models.Comuna{
		{
			ID: 1,
			Nombre: "Concepción",
			CiudadID: 1,
		},
		{
			ID: 2,
			Nombre: "Coronel",
			CiudadID: 1,
		},
		{
			ID: 3,
			Nombre: "Chiguayante",
			CiudadID: 1,
		},
		{
			ID: 4,
			Nombre: "Florida",
			CiudadID: 1,
		},
		{
			ID: 5,
			Nombre: "Hualpén",
			CiudadID: 1,
		},
		{
			ID: 6,
			Nombre: "Hualqui",
			CiudadID: 1,
		},
		{
			ID: 7,
			Nombre: "Lota",
			CiudadID: 1,
		},
		{
			ID: 8,
			Nombre: "Penco",
			CiudadID: 1,
		},
		{
			ID: 9,
			Nombre: "San Pedro de la Paz",
			CiudadID: 1,
		},
		{
			ID: 10,
			Nombre: "Santa Juana",
			CiudadID: 1,
		},
		{
			ID: 11,
			Nombre: "Tomé",
			CiudadID: 1,
		},
		{
			ID: 12,
			Nombre: "Talcahuano",
			CiudadID: 3,
		},
		{
			ID: 13,
			Nombre: "Chillán",
			CiudadID: 2,
		},
		{
			ID: 14,
			Nombre: "Chillán Viejo",
			CiudadID: 2,
		},
	}

	dias := []models.Dia{
		{Nombre: "Lunes"},
		{Nombre: "Martes"},
		{Nombre: "Miércoles"},
		{Nombre: "Jueves"},
		{Nombre: "Viernes"},
		{Nombre: "Sábado"},
		{Nombre: "Domingo"},
	}

	horas := []models.Hora{
		{Hora: "00:00"},
		{Hora: "01:00"},
		{Hora: "02:00"},
		{Hora: "03:00"},
		{Hora: "04:00"},
		{Hora: "05:00"},
		{Hora: "06:00"},
		{Hora: "07:00"},
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
		{Hora: "23:00"}}

	gdb.Create(&pais)
	gdb.Create(&regiones)
	gdb.Create(&ciudades)
	gdb.Create(&comunas)
	gdb.Create(&dias)
	gdb.Create(&horas)
}

func Close(gdb *gorm.DB) {
	sqlDB, err := gdb.DB()
	if err != nil {
		panic("failed to get gorm db")
	}
	sqlDB.Close()
}
