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

func repartidorData() (models.Repartidor, *sqlmock.Rows) {
	repartidor := models.Repartidor{
		RUT:               "19120771-8",
		Contrasena:        "contraseñaFalsa1234",
		Email:             "email@falso.cl",
		Nombres:           "Juan Antonio",
		ApellidoPaterno:   "Kotton",
		ApellidoMaterno:   "Ciruela",
		Telefono:          99834553,
		Direccion:         "Calle falsa 123",
		FechaNacimiento:   utils.DateTimeNow(),
		FechaRegistracion: utils.DateTimeNow(),
		TipoLicencia:      2,
		FechaLicencia:     utils.DateTimeNow(),
	}

	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "rut", "contrasena", "email", "nombres", "apellido_paterno", "apellido_materno", "telefono", "direccion", "fecha_nacimiento", "fecha_registracion", "tipo_licencia", "fecha_licencia"}).
		AddRow("1", utils.DateTimeNow(), utils.DateTimeNow(), nil, repartidor.RUT, repartidor.Contrasena, repartidor.Email, repartidor.Nombres, repartidor.ApellidoPaterno, repartidor.ApellidoMaterno, repartidor.Telefono, repartidor.Direccion, repartidor.FechaNacimiento, repartidor.FechaRegistracion, repartidor.TipoLicencia, repartidor.FechaLicencia).
		AddRow("2", utils.DateTimeNow(), utils.DateTimeNow(), nil, repartidor.RUT+"2", repartidor.Contrasena+"2", repartidor.Email+"2", repartidor.Nombres+"2", repartidor.ApellidoPaterno+"2", repartidor.ApellidoMaterno+"2", repartidor.Telefono+2, repartidor.Direccion+"2", repartidor.FechaNacimiento, repartidor.FechaRegistracion, 1, repartidor.FechaLicencia)

	return repartidor, rows
}

func setupRepartidor(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error `%s` was not expected when opening a stub database connection", err)
	}
	if sqlDB == nil {
		t.Error("mock db is null")
	}
	if mock == nil {
		t.Error("sqlmock is null")
	}

	gdb, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "sqlmock_db_0",
		DriverName:                "mysql",
		Conn:                      sqlDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		t.Errorf("Failed to open gorm v2 db, got error: %v", err)
	}
	return gdb, mock
}

func teardownRepartidor(gdb *gorm.DB) {
	database.Close(gdb)
}

func TestPostRepartidor(t *testing.T) {
	gdb, mock := setupRepartidor(t)
	repartidor, _ := repartidorData()

	mock.ExpectBegin()

	mock.ExpectExec(
		regexp.QuoteMeta("INSERT INTO `repartidor` (`created_at`,`updated_at`,`deleted_at`,`rut`,`contrasena`,`email`,`nombres`,`apellido_paterno`,`apellido_materno`,`telefono`,`direccion`,`fecha_nacimiento`,`fecha_registracion`,`tipo_licencia`,`fecha_licencia`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")).
		WithArgs(utils.AnyTime{}, utils.AnyTime{}, nil, repartidor.RUT, repartidor.Contrasena, repartidor.Email, repartidor.Nombres, repartidor.ApellidoPaterno, repartidor.ApellidoMaterno, repartidor.Telefono, repartidor.Direccion, repartidor.FechaNacimiento, repartidor.FechaRegistracion, repartidor.TipoLicencia, repartidor.FechaLicencia).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	postRepartidor(repartidor, gdb)

	err := mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Failed to meet expectations, got error: %v", err)
	}

	teardownRepartidor(gdb)
}

func TestListRepartidores(t *testing.T) {
	gdb, mock := setupRepartidor(t)

	_, rows := repartidorData()

	mock.ExpectQuery(
		regexp.QuoteMeta("SELECT * FROM `repartidor` WHERE `repartidor`.`deleted_at` IS NULL ORDER BY ID asc")).
		WillReturnRows(rows)

	var Repartidores []models.Repartidor
	listRepartidores(&Repartidores, gdb)

	err := mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Failed to meet expectations, got error: %v", err)
	}

	teardownRepartidor(gdb)
}

func TestGetRepartidors(t *testing.T) {
	gdb, mock := setupRepartidor(t)

	repartidor, rows := repartidorData()

	mock.ExpectQuery(
		regexp.QuoteMeta("SELECT * FROM `repartidor` WHERE `repartidor`.`id` = ? AND `repartidor`.`deleted_at` IS NULL ORDER BY ID asc")).
		WithArgs("1").
		WillReturnRows(rows)

	result := getRepartidor(&repartidor, "1", gdb)
	if result.Error != nil {
		t.Errorf("getRepartidor failed: %v", result.Error)
	} else if result.RowsAffected == 0 {
		t.Errorf("getRepartidor had 0 rows")
	}

	err := mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Failed to meet expectations, got error: %v", err)
	}

	teardownRepartidor(gdb)
}

func TestPutRepartidors(t *testing.T) {
	gdb, mock := setupRepartidor(t)

	repartidor, _ := repartidorData()
	repartidor2 := repartidor
	repartidor2.ID = 1
	repartidor3 := repartidor
	repartidor3.ID = 3

	mock.ExpectBegin()
	mock.ExpectExec(
		regexp.QuoteMeta("INSERT INTO `repartidor` (`created_at`,`updated_at`,`deleted_at`,`rut`,`contrasena`,`email`,`nombres`,`apellido_paterno`,`apellido_materno`,`telefono`,`direccion`,`fecha_nacimiento`,`fecha_registracion`,`tipo_licencia`,`fecha_licencia`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")).
		WithArgs(utils.AnyTime{}, utils.AnyTime{}, nil, repartidor.RUT, repartidor.Contrasena, repartidor.Email, repartidor.Nombres, repartidor.ApellidoPaterno, repartidor.ApellidoMaterno, repartidor.Telefono, repartidor.Direccion, repartidor.FechaNacimiento, repartidor.FechaRegistracion, repartidor.TipoLicencia, repartidor.FechaLicencia).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	mock.ExpectBegin()
	mock.ExpectExec(
		regexp.QuoteMeta("UPDATE `repartidor` SET `created_at`=?,`updated_at`=?,`deleted_at`=?,`rut`=?,`contrasena`=?,`email`=?,`nombres`=?,`apellido_paterno`=?,`apellido_materno`=?,`telefono`=?,`direccion`=?,`fecha_nacimiento`=?,`fecha_registracion`=?,`tipo_licencia`=?,`fecha_licencia`=? WHERE `repartidor`.`deleted_at` IS NULL AND `id` = ?")).
		WithArgs(utils.AnyTime{}, utils.AnyTime{}, nil, repartidor2.RUT, repartidor2.Contrasena, repartidor2.Email, repartidor2.Nombres, repartidor2.ApellidoPaterno, repartidor2.ApellidoMaterno, repartidor2.Telefono, repartidor2.Direccion, repartidor2.FechaNacimiento, repartidor2.FechaRegistracion, repartidor2.TipoLicencia, repartidor2.FechaLicencia, 1).
		WillReturnResult(sqlmock.NewResult(2, 1))
	mock.ExpectCommit()

	mock.ExpectBegin()
	mock.ExpectExec(
		regexp.QuoteMeta("UPDATE `repartidor` SET `created_at`=?,`updated_at`=?,`deleted_at`=?,`rut`=?,`contrasena`=?,`email`=?,`nombres`=?,`apellido_paterno`=?,`apellido_materno`=?,`telefono`=?,`direccion`=?,`fecha_nacimiento`=?,`fecha_registracion`=?,`tipo_licencia`=?,`fecha_licencia`=? WHERE `repartidor`.`deleted_at` IS NULL AND `id` = ?")).
		WithArgs(utils.AnyTime{}, utils.AnyTime{}, nil, repartidor3.RUT, repartidor3.Contrasena, repartidor3.Email, repartidor3.Nombres, repartidor3.ApellidoPaterno, repartidor3.ApellidoMaterno, repartidor3.Telefono, repartidor3.Direccion, repartidor3.FechaNacimiento, repartidor3.FechaRegistracion, repartidor3.TipoLicencia, repartidor3.FechaLicencia, 3).
		WillReturnResult(sqlmock.NewResult(3, 0))
	mock.ExpectCommit()

	mock.ExpectQuery(
		regexp.QuoteMeta("SELECT * FROM `repartidor` WHERE `repartidor`.`deleted_at` IS NULL AND `id` = ? LIMIT 1")).
		WithArgs(repartidor3.ID).
		WillReturnRows(mock.NewRows(nil))

	mock.ExpectBegin()
	mock.ExpectExec(
		regexp.QuoteMeta("INSERT INTO `repartidor` (`created_at`,`updated_at`,`deleted_at`,`rut`,`contrasena`,`email`,`nombres`,`apellido_paterno`,`apellido_materno`,`telefono`,`direccion`,`fecha_nacimiento`,`fecha_registracion`,`tipo_licencia`,`fecha_licencia`,`id`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")).
		WithArgs(utils.AnyTime{}, utils.AnyTime{}, nil, repartidor3.RUT, repartidor3.Contrasena, repartidor3.Email, repartidor3.Nombres, repartidor3.ApellidoPaterno, repartidor3.ApellidoMaterno, repartidor3.Telefono, repartidor3.Direccion, repartidor3.FechaNacimiento, repartidor3.FechaRegistracion, repartidor3.TipoLicencia, repartidor3.FechaLicencia, 3).
		WillReturnResult(sqlmock.NewResult(3, 1))
	mock.ExpectCommit()

	result := putRepartidor(&repartidor, gdb)
	if result.Error != nil {
		t.Errorf("putRepartidor failed: %v", result.Error)
	} else if result.RowsAffected == 0 {
		t.Errorf("putRepartidor had 0 rows")
	}

	result = putRepartidor(&repartidor2, gdb)
	if result.Error != nil {
		t.Errorf("putRepartidor failed: %v", result.Error)
	} else if result.RowsAffected == 0 {
		t.Errorf("putRepartidor had 0 rows")
	}

	result = putRepartidor(&repartidor3, gdb)
	if result.Error != nil {
		t.Errorf("putRepartidor failed: %v", result.Error)
	} else if result.RowsAffected == 0 {
		t.Errorf("putRepartidor had 0 rows")
	}

	err := mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Failed to meet expectations, got error: %v", err)
	}

	teardownRepartidor(gdb)
}

func TestPatchRepartidors(t *testing.T) {
	gdb, mock := setupRepartidor(t)

	repartidor := models.Repartidor{
		Nombres: "Federico Casanova",
	}
	repartidor.ID = 1

	mock.ExpectBegin()
	mock.ExpectExec(
		regexp.QuoteMeta("UPDATE `repartidor` SET `updated_at`=?,`nombres`=? WHERE `repartidor`.`deleted_at` IS NULL AND `id` = ?")).
		WithArgs(utils.AnyTime{}, repartidor.Nombres, repartidor.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	result := patchRepartidor(&repartidor, gdb)
	if result.Error != nil {
		t.Errorf("patchRepartidor failed: %v", result.Error)
	} else if result.RowsAffected == 0 {
		t.Errorf("patchRepartidor had 0 rows")
	}

	err := mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Failed to meet expectations, got error: %v", err)
	}

	teardownRepartidor(gdb)
}

func TestDeleteRepartidors(t *testing.T) {
	gdb, mock := setupRepartidor(t)

	repartidor, _ := repartidorData()
	repartidor.ID = 1

	mock.ExpectBegin()
	mock.ExpectExec(
		regexp.QuoteMeta("DELETE FROM `repartidor` WHERE `repartidor`.`id` = ?")).
		WithArgs(repartidor.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	mock.ExpectBegin()
	mock.ExpectExec(
		regexp.QuoteMeta("UPDATE `repartidor` SET `deleted_at`=? WHERE `repartidor`.`id` = ? AND `repartidor`.`deleted_at` IS NULL")).
		WithArgs(utils.AnyTime{}, repartidor.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	result := deleteRepartidor(&repartidor, true, gdb)
	if result.Error != nil {
		t.Errorf("deleteRepartidor failed: %v", result.Error)
	} else if result.RowsAffected == 0 {
		t.Errorf("deleteRepartidor had 0 rows")
	}

	result = deleteRepartidor(&repartidor, false, gdb)
	if result.Error != nil {
		t.Errorf("deleteRepartidor failed: %v", result.Error)
	} else if result.RowsAffected == 0 {
		t.Errorf("deleteRepartidor had 0 rows")
	}

	err := mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Failed to meet expectations, got error: %v", err)
	}

	teardownRepartidor(gdb)
}
