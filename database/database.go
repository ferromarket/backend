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
	gdb.AutoMigrate(&models.Producto{})
	gdb.AutoMigrate(&models.Categoria{})
	gdb.AutoMigrate(&models.Especificacion{})
	gdb.AutoMigrate(&models.EspecificacionData{})
	gdb.AutoMigrate(&models.EspecificacionNombre{})
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
		&models.Usuario{},
		&models.Producto{},
		&models.Categoria{},
		&models.Especificacion{},
		&models.EspecificacionData{},
		&models.EspecificacionNombre{})
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
		{ID: 1, Nombre: "Lunes"},
		{ID: 2, Nombre: "Martes"},
		{ID: 3, Nombre: "Miércoles"},
		{ID: 4, Nombre: "Jueves"},
		{ID: 5, Nombre: "Viernes"},
		{ID: 6, Nombre: "Sábado"},
		{ID: 7, Nombre: "Domingo"},
	}

	horas := []models.Hora{
		{ID: 1, Hora: "00:00"},
		{ID: 2, Hora: "01:00"},
		{ID: 3, Hora: "02:00"},
		{ID: 4, Hora: "03:00"},
		{ID: 5, Hora: "04:00"},
		{ID: 6, Hora: "05:00"},
		{ID: 7, Hora: "06:00"},
		{ID: 8, Hora: "07:00"},
		{ID: 9, Hora: "08:00"},
		{ID: 10, Hora: "09:00"},
		{ID: 11, Hora: "10:00"},
		{ID: 12, Hora: "11:00"},
		{ID: 13, Hora: "12:00"},
		{ID: 14, Hora: "13:00"},
		{ID: 15, Hora: "14:00"},
		{ID: 16, Hora: "15:00"},
		{ID: 17, Hora: "16:00"},
		{ID: 18, Hora: "17:00"},
		{ID: 19, Hora: "18:00"},
		{ID: 20, Hora: "19:00"},
		{ID: 21, Hora: "20:00"},
		{ID: 22, Hora: "21:00"},
		{ID: 23, Hora: "22:00"},
		{ID: 24, Hora: "23:00"}}

	gdb.Create(&pais)
	gdb.Create(&regiones)
	gdb.Create(&ciudades)
	gdb.Create(&comunas)
	gdb.Create(&dias)
	gdb.Create(&horas)

	categorias := []models.Categoria{
		{
			ID:     1,
			Nombre: "Herramientas Manuales",
		},
		{
			ID:     2,
			Nombre: "Herramientas Electricas",
		},
		{
			ID:     3,
			Nombre: "Herramientas de Medicion",
		},
		{
			ID:     4,
			Nombre: "Herramientas para Jardin",
		},
		{
			ID:     5,
			Nombre: "Herramientas Industriales",
		},
		{
			ID:     6,
			Nombre: "Construccion",
		},
	}

	gdb.Create(&categorias)

	var categoriaID *uint64 = new(uint64)

	*categoriaID = 1
	categoria := models.Categoria{
		CategoriaID: categoriaID,
		Nombre:      "Hijo",
	}
	gdb.Create(&categoria)
	*categoriaID = 7
	categoria = models.Categoria{
		CategoriaID: categoriaID,
		Nombre:      "Hijo2",
	}
	gdb.Create(&categoria)
}

func Close(gdb *gorm.DB) {
	sqlDB, err := gdb.DB()
	if err != nil {
		panic("failed to get gorm db")
	}
	sqlDB.Close()
}
