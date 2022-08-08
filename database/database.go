package database

import (
	"os"

	"github.com/ferromarket/backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Connect() *gorm.DB {
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
		DSN:               dbUser + ":" + dbPass + "@tcp(" + dbHost + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local",
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
	ferreteria := models.Ferreteria{
		Nombre:      "Chris's Hardware Store",
		Direccion:   "Canto del Valle 1777",
		Descripcion: "Buy something here!",
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

	horarios := []models.FerreteriaHorario{
		{
			FerreteriaID: 1,
			DiaID:        1,
			AbrirID:      9,
			CerrarID:     22,
		},
	}

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

	gdb.Create(&ferreteria)
	gdb.Create(&dias)
	gdb.Create(&horas)
	gdb.Create(&horarios)

}

func Close(gdb *gorm.DB) {
	sqlDB, err := gdb.DB()
	if err != nil {
		panic("failed to get gorm db")
	}
	sqlDB.Close()
}
