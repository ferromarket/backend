package controllers

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ferromarket/backend/database"
	"github.com/ferromarket/backend/models"
	"github.com/ferromarket/backend/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func vehiculoData() (models.Vehiculo, *sqlmock.Rows) {
	vehiculo := models.Vehiculo{
		RepartidorID: 1,
		Patente: "SA23-233SS",
		Marca: "Audi",
		Modelo: "A3",
		Ano: 2003,
		PermisoCirculacion: utils.DateNow(),
		Seguro: utils.DateNow(),
	}

	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "repartidor_id", "patente", "marca", "modelo", "ano", "permiso_circulacion", "seguro"}).
		AddRow("1", utils.DateNow(), utils.DateNow(), nil, vehiculo.RepartidorID, vehiculo.Patente, vehiculo.Marca, vehiculo.Modelo, vehiculo.Ano, vehiculo.PermisoCirculacion, vehiculo.Seguro).
		AddRow("2", utils.DateNow(), utils.DateNow(), nil, vehiculo.RepartidorID, vehiculo.Patente + "2", vehiculo.Marca + "2", vehiculo.Modelo + "2", vehiculo.Ano+2, vehiculo.PermisoCirculacion, vehiculo.Seguro)

	return vehiculo, rows
}

func setupVehiculo(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	sqlDB, mock, err := sqlmock.New()
	if (err != nil) {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	if (sqlDB == nil) {
		t.Error("mock db is null")
	}
	if (mock == nil) {
		t.Error("sqlmock is null")
	}

	gdb, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "sqlmock_db_0",
		DriverName: "mysql",
		Conn: sqlDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if (err != nil) {
		t.Errorf("Failed to open gorm v2 db, got error: %v", err)
	}
	return gdb, mock
}

func teardownVehiculo(gdb *gorm.DB) {
	database.Close(gdb)
}

func TestPostVehiculo(t *testing.T) {
	gdb, mock := setupVehiculo(t)
	vehiculo, _ := vehiculoData()

	mock.ExpectBegin()

	mock.ExpectExec(
		regexp.QuoteMeta("INSERT INTO `vehiculo` (`created_at`,`updated_at`,`deleted_at`,`repartidor_id`,`patente`,`marca`,`modelo`,`ano`,`permiso_circulacion`,`seguro`) VALUES (?,?,?,?,?,?,?,?,?,?)")).
		WithArgs(utils.AnyTime{}, utils.AnyTime{}, nil, vehiculo.RepartidorID, vehiculo.Patente, vehiculo.Marca, vehiculo.Modelo, vehiculo.Ano, vehiculo.PermisoCirculacion, vehiculo.Seguro).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	postVehiculo(vehiculo, gdb)

	err := mock.ExpectationsWereMet()
	if (err != nil) {
		t.Errorf("Failed to meet expectations, got error: %v", err)
	}

	teardownVehiculo(gdb)
}

func TestListVehiculos(t *testing.T) {
	gdb, mock := setupVehiculo(t)

	_, rows := vehiculoData()

	mock.ExpectQuery(
		regexp.QuoteMeta("SELECT * FROM `vehiculo` WHERE `vehiculo`.`deleted_at` IS NULL ORDER BY ID asc")).
		WillReturnRows(rows)

	mock.ExpectQuery(
		regexp.QuoteMeta("SELECT * FROM `repartidor` WHERE `repartidor`.`id` = ? AND `repartidor`.`deleted_at` IS NULL")).
		WithArgs(1).
		WillReturnRows(mock.NewRows([]string{"id"}).
			AddRow(1))

	var vehiculos []models.Vehiculo
	listVehiculos(&vehiculos, gdb)

	err := mock.ExpectationsWereMet()
	if (err != nil) {
		t.Errorf("Failed to meet expectations, got error: %v", err)
	}

	teardownVehiculo(gdb)
}

func TestGetVehiculo(t *testing.T) {
	gdb, mock := setupVehiculo(t)

	vehiculo, rows := vehiculoData()

	mock.ExpectQuery(
		regexp.QuoteMeta("SELECT * FROM `vehiculo` WHERE `vehiculo`.`id` = ? AND `vehiculo`.`deleted_at` IS NULL ORDER BY ID asc")).
		WithArgs("1").
		WillReturnRows(rows)

	mock.ExpectQuery(
		regexp.QuoteMeta("SELECT * FROM `repartidor` WHERE `repartidor`.`id` = ? AND `repartidor`.`deleted_at` IS NULL")).
		WithArgs(1).
		WillReturnRows(mock.NewRows(nil))

	result := getVehiculo(&vehiculo, "1", gdb)
	if (result.Error != nil) {
		t.Errorf("Vehiculo failed: %v", result.Error)
	} else if (result.RowsAffected == 0) {
		t.Errorf("Vehiculo had 0 rows")
	}

	err := mock.ExpectationsWereMet()
	if (err != nil) {
		t.Errorf("Failed to meet expectations, got error: %v", err)
	}

	teardownVehiculo(gdb)
}

func TestPutVehiculos(t *testing.T) {
	gdb, mock := setupVehiculo(t)

	vehiculo, _ := vehiculoData()
	vehiculo2 := vehiculo
	vehiculo2.ID = 1
	vehiculo3 := vehiculo
	vehiculo3.ID = 3

	mock.ExpectBegin()
	mock.ExpectExec(
		regexp.QuoteMeta("INSERT INTO `vehiculo` (`created_at`,`updated_at`,`deleted_at`,`repartidor_id`,`patente`,`marca`,`modelo`,`ano`,`permiso_circulacion`,`seguro`) VALUES (?,?,?,?,?,?,?,?,?,?)")).
		WithArgs(utils.AnyTime{}, utils.AnyTime{}, nil, vehiculo.RepartidorID, vehiculo.Patente, vehiculo.Marca, vehiculo.Modelo, vehiculo.Ano, vehiculo.PermisoCirculacion, vehiculo.Seguro).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	mock.ExpectBegin()
	mock.ExpectExec(
		regexp.QuoteMeta("UPDATE `vehiculo` SET `created_at`=?,`updated_at`=?,`deleted_at`=?,`repartidor_id`=?,`patente`=?,`marca`=?,`modelo`=?,`ano`=?,`permiso_circulacion`=?,`seguro`=?  WHERE `vehiculo`.`deleted_at` IS NULL AND `id` = ?")).
		WithArgs(utils.AnyTime{}, utils.AnyTime{}, nil, vehiculo2.RepartidorID, vehiculo2.Patente, vehiculo2.Marca, vehiculo2.Modelo, vehiculo2.Ano, vehiculo2.PermisoCirculacion, vehiculo2.Seguro, 1).
		WillReturnResult(sqlmock.NewResult(2, 1))
	mock.ExpectCommit()

	mock.ExpectBegin()
	mock.ExpectExec(
		regexp.QuoteMeta("UPDATE `vehiculo` SET `created_at`=?,`updated_at`=?,`deleted_at`=?,`repartidor_id`=?,`patente`=?,`marca`=?,`modelo`=?,`ano`=?,`permiso_circulacion`=?,`seguro`=? WHERE `vehiculo`.`deleted_at` IS NULL AND `id` = ?")).
		WithArgs(utils.AnyTime{}, utils.AnyTime{}, nil, vehiculo3.RepartidorID, vehiculo3.Patente, vehiculo3.Marca, vehiculo3.Modelo, vehiculo3.Ano, vehiculo3.PermisoCirculacion, vehiculo3.Seguro, 3).
		WillReturnResult(sqlmock.NewResult(3, 0))
	mock.ExpectCommit()

	mock.ExpectQuery(
		regexp.QuoteMeta("SELECT * FROM `vehiculo` WHERE `vehiculo`.`deleted_at` IS NULL AND `id` = ? LIMIT 1")).
		WithArgs(vehiculo3.ID).
		WillReturnRows(mock.NewRows(nil))

	mock.ExpectBegin()
	mock.ExpectExec(
		regexp.QuoteMeta("INSERT INTO `vehiculo` (`created_at`,`updated_at`,`deleted_at`,`repartidor_id`,`patente`,`marca`,`modelo`,`ano`,`permiso_circulacion`,`seguro`,`id`) VALUES (?,?,?,?,?,?,?,?,?,?,?)")).
		WithArgs(utils.AnyTime{}, utils.AnyTime{}, nil, vehiculo3.RepartidorID, vehiculo3.Patente, vehiculo3.Marca, vehiculo3.Modelo, vehiculo3.Ano, vehiculo3.PermisoCirculacion, vehiculo3.Seguro, 3).
		WillReturnResult(sqlmock.NewResult(3, 1))
	mock.ExpectCommit()

	result := putVehiculo(&vehiculo, gdb)
	if (result.Error != nil) {
		t.Errorf("putVehiculo failed: %v", result.Error)
	} else if (result.RowsAffected == 0) {
		t.Errorf("putVehiculo had 0 rows")
	}

	result = putVehiculo(&vehiculo2, gdb)
	if (result.Error != nil) {
		t.Errorf("putVehiculo failed: %v", result.Error)
	} else if (result.RowsAffected == 0) {
		t.Errorf("putVehiculo had 0 rows")
	}

	result = putVehiculo(&vehiculo3, gdb)
	if (result.Error != nil) {
		t.Errorf("putVehiculo failed: %v", result.Error)
	} else if (result.RowsAffected == 0) {
		t.Errorf("putVehiculo had 0 rows")
	}

	err := mock.ExpectationsWereMet()
	if (err != nil) {
		t.Errorf("Failed to meet expectations, got error: %v", err)
	}

	teardownVehiculo(gdb)
}

func TestPatchVehiculos(t *testing.T) {
	gdb, mock := setupVehiculo(t)

	vehiculo := models.Vehiculo{
		Patente: "123kk2-123wwq",
	}
	vehiculo.ID = 1

	mock.ExpectBegin()
	mock.ExpectExec(
		regexp.QuoteMeta("UPDATE `vehiculo` SET `updated_at`=?,`patente`=? WHERE `vehiculo`.`deleted_at` IS NULL AND `id` = ?")).
		WithArgs(utils.AnyTime{}, vehiculo.Patente, vehiculo.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	result := patchVehiculo(&vehiculo, gdb)
	if (result.Error != nil) {
		t.Errorf("patchVehiculo failed: %v", result.Error)
	} else if (result.RowsAffected == 0) {
		t.Errorf("patchVehiculo had 0 rows")
	}

	err := mock.ExpectationsWereMet()
	if (err != nil) {
		t.Errorf("Failed to meet expectations, got error: %v", err)
	}

	teardownVehiculo(gdb)
}

func TestDeleteVehiculo(t *testing.T) {
	gdb, mock := setupVehiculo(t)

	vehiculo, _ := vehiculoData()
	vehiculo.ID = 1

	mock.ExpectBegin()
	mock.ExpectExec(
		regexp.QuoteMeta("DELETE FROM `vehiculo` WHERE `vehiculo`.`id` = ?")).
		WithArgs(vehiculo.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	mock.ExpectBegin()
	mock.ExpectExec(
		regexp.QuoteMeta("UPDATE `vehiculo` SET `deleted_at`=? WHERE `vehiculo`.`id` = ? AND `vehiculo`.`deleted_at` IS NULL")).
		WithArgs(utils.AnyTime{}, vehiculo.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	result := deleteVehiculo(&vehiculo, true, gdb)
	if (result.Error != nil) {
		t.Errorf("deleteVehiculo failed: %v", result.Error)
	} else if (result.RowsAffected == 0) {
		t.Errorf("deleteVehiculo had 0 rows")
	}

	result = deleteVehiculo(&vehiculo, false, gdb)
	if (result.Error != nil) {
		t.Errorf("deleteVehiculo failed: %v", result.Error)
	} else if (result.RowsAffected == 0) {
		t.Errorf("deleteVehiculo had 0 rows")
	}

	err := mock.ExpectationsWereMet()
	if (err != nil) {
		t.Errorf("Failed to meet expectations, got error: %v", err)
	}

	teardownVehiculo(gdb)
}
