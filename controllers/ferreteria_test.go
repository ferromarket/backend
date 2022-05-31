package controllers

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ferromarket/backend/database"
	"github.com/ferromarket/backend/models"
	"github.com/ferromarket/backend/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func ferreteriaData() (models.Ferreteria, *sqlmock.Rows) {
	ferreteria := models.Ferreteria{
		Nombre: "Chris's hardware",
		ComunaID: 1,
		Direccion: "214 Edison Dr.",
		Descripcion: "blah blah blah",
	}

	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "nombre", "direccion", "descripcion", "comuna_id"}).
		AddRow("1", time.Time{}, time.Time{}, nil, ferreteria.Nombre, ferreteria.Direccion, ferreteria.Descripcion, 1).
		AddRow("2", time.Time{}, time.Time{}, nil, ferreteria.Nombre + "2", ferreteria.Direccion + "2", ferreteria.Descripcion + "2", 2)

	return ferreteria, rows
}

func setup(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
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

func teardown(gdb *gorm.DB) {
	database.Close(gdb)
}

func TestPostFerreteria(t *testing.T) {
	gdb, mock := setup(t)

	ferreteria := models.Ferreteria{
		Nombre: "Chris's hardware",
		ComunaID: 1,
		Direccion: "214 Edison Dr.",
		Descripcion: "blah blah blah",
	}

	mock.ExpectBegin()

	mock.ExpectExec(
		regexp.QuoteMeta("INSERT INTO `ferreteria` (`created_at`,`updated_at`,`deleted_at`,`nombre`,`comuna_id`,`direccion`,`descripcion`) VALUES (?,?,?,?,?,?,?)")).
		WithArgs(utils.AnyTime{}, utils.AnyTime{}, nil, ferreteria.Nombre, ferreteria.ComunaID, ferreteria.Direccion, ferreteria.Descripcion).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	postFerreteria(ferreteria, gdb)

	err := mock.ExpectationsWereMet()
	if (err != nil) {
		t.Errorf("Failed to meet expectations, got error: %v", err)
	}

	teardown(gdb)
}

func TestListFerreterias(t *testing.T) {
	gdb, mock := setup(t)

	_, rows := ferreteriaData()

	mock.ExpectQuery(
		regexp.QuoteMeta("SELECT `ferreteria`.`id`,`ferreteria`.`created_at`,`ferreteria`.`updated_at`,`ferreteria`.`deleted_at`,`ferreteria`.`nombre`,`ferreteria`.`comuna_id`,`ferreteria`.`direccion`,`ferreteria`.`descripcion` FROM `ferreteria` LEFT JOIN ferreteria_horario fh ON ferreteria.id = fh.ferreteria_id WHERE `ferreteria`.`deleted_at` IS NULL ORDER BY ID asc")).
		WillReturnRows(rows)

	mock.ExpectQuery(
		regexp.QuoteMeta("SELECT * FROM `comuna` WHERE `comuna`.`id` IN (?,?) AND `comuna`.`deleted_at` IS NULL")).
		WithArgs(1, 2).
		WillReturnRows(mock.NewRows([]string{"id"}).
			AddRow(1))

	mock.ExpectQuery(
		regexp.QuoteMeta("SELECT * FROM `ferreteria_horario` WHERE `ferreteria_horario`.`ferreteria_id` IN (?,?) AND `ferreteria_horario`.`deleted_at` IS NULL")).
		WithArgs(1, 2).
		WillReturnRows(mock.NewRows([]string{"ferreteria_id"}).
			AddRow(1))

	var ferreterias []models.Ferreteria
	listFerreterias(&ferreterias, gdb)

	err := mock.ExpectationsWereMet()
	if (err != nil) {
		t.Errorf("Failed to meet expectations, got error: %v", err)
	}

	teardown(gdb)
}

func TestGetFerreterias(t *testing.T) {
	gdb, mock := setup(t)

	ferreteria, rows := ferreteriaData()

	mock.ExpectQuery(
		regexp.QuoteMeta("SELECT `ferreteria`.`id`,`ferreteria`.`created_at`,`ferreteria`.`updated_at`,`ferreteria`.`deleted_at`,`ferreteria`.`nombre`,`ferreteria`.`comuna_id`,`ferreteria`.`direccion`,`ferreteria`.`descripcion` FROM `ferreteria` LEFT JOIN ferreteria_horario fh ON ferreteria.id = fh.ferreteria_id WHERE `ferreteria`.`id` = ? AND `ferreteria`.`deleted_at` IS NULL ORDER BY ID asc")).
		WithArgs("1").
		WillReturnRows(rows)

	mock.ExpectQuery(
		regexp.QuoteMeta("SELECT * FROM `comuna` WHERE `comuna`.`id` = ? AND `comuna`.`deleted_at` IS NULL")).
		WithArgs(1).
		WillReturnRows(mock.NewRows(nil))

	mock.ExpectQuery(
		regexp.QuoteMeta("SELECT * FROM `ferreteria_horario` WHERE `ferreteria_horario`.`ferreteria_id` = ? AND `ferreteria_horario`.`deleted_at` IS NULL")).
		WithArgs(1).
		WillReturnRows(mock.NewRows(nil))

	result := getFerreteria(&ferreteria, "1", gdb)
	if (result.Error != nil) {
		t.Errorf("getFerreteria failed: %v", result.Error)
	} else if (result.RowsAffected == 0) {
		t.Errorf("getFerreteria had 0 rows")
	}

	err := mock.ExpectationsWereMet()
	if (err != nil) {
		t.Errorf("Failed to meet expectations, got error: %v", err)
	}

	teardown(gdb)
}

func TestPutFerreterias(t *testing.T) {
	gdb, mock := setup(t)

	ferreteria, _ := ferreteriaData()
	ferreteria2 := ferreteria
	ferreteria2.ID = 1
	ferreteria3 := ferreteria
	ferreteria3.ID = 3

	mock.ExpectBegin()
	mock.ExpectExec(
		regexp.QuoteMeta("INSERT INTO `ferreteria` (`created_at`,`updated_at`,`deleted_at`,`nombre`,`comuna_id`,`direccion`,`descripcion`) VALUES (?,?,?,?,?,?,?)")).
		WithArgs(utils.AnyTime{}, utils.AnyTime{}, nil, ferreteria.Nombre, ferreteria.ComunaID, ferreteria.Direccion, ferreteria.Descripcion).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	mock.ExpectBegin()
	mock.ExpectExec(
		regexp.QuoteMeta("UPDATE `ferreteria` SET `created_at`=?,`updated_at`=?,`deleted_at`=?,`nombre`=?,`comuna_id`=?,`direccion`=?,`descripcion`=? WHERE `ferreteria`.`deleted_at` IS NULL AND `id` = ?")).
		WithArgs(utils.AnyTime{}, utils.AnyTime{}, nil, ferreteria2.Nombre, ferreteria2.ComunaID, ferreteria2.Direccion, ferreteria2.Descripcion, 1).
		WillReturnResult(sqlmock.NewResult(2, 1))
	mock.ExpectCommit()

	mock.ExpectBegin()
	mock.ExpectExec(
		regexp.QuoteMeta("UPDATE `ferreteria` SET `created_at`=?,`updated_at`=?,`deleted_at`=?,`nombre`=?,`comuna_id`=?,`direccion`=?,`descripcion`=? WHERE `ferreteria`.`deleted_at` IS NULL AND `id` = ?")).
		WithArgs(utils.AnyTime{}, utils.AnyTime{}, nil, ferreteria3.Nombre, ferreteria3.ComunaID, ferreteria3.Direccion, ferreteria3.Descripcion, 3).
		WillReturnResult(sqlmock.NewResult(3, 0))
	mock.ExpectCommit()

	mock.ExpectQuery(
		regexp.QuoteMeta("SELECT * FROM `ferreteria` WHERE `ferreteria`.`deleted_at` IS NULL AND `id` = ? LIMIT 1")).
		WithArgs(ferreteria3.ID).
		WillReturnRows(mock.NewRows(nil))

	mock.ExpectBegin()
	mock.ExpectExec(
		regexp.QuoteMeta("INSERT INTO `ferreteria` (`created_at`,`updated_at`,`deleted_at`,`nombre`,`comuna_id`,`direccion`,`descripcion`,`id`) VALUES (?,?,?,?,?,?,?,?)")).
		WithArgs(utils.AnyTime{}, utils.AnyTime{}, nil, ferreteria3.Nombre, ferreteria3.ComunaID, ferreteria3.Direccion, ferreteria3.Descripcion, 3).
		WillReturnResult(sqlmock.NewResult(3, 1))
	mock.ExpectCommit()

	result := putFerreteria(&ferreteria, gdb)
	if (result.Error != nil) {
		t.Errorf("putFerreteria failed: %v", result.Error)
	} else if (result.RowsAffected == 0) {
		t.Errorf("putFerreteria had 0 rows")
	}

	result = putFerreteria(&ferreteria2, gdb)
	if (result.Error != nil) {
		t.Errorf("putFerreteria failed: %v", result.Error)
	} else if (result.RowsAffected == 0) {
		t.Errorf("putFerreteria had 0 rows")
	}

	result = putFerreteria(&ferreteria3, gdb)
	if (result.Error != nil) {
		t.Errorf("putFerreteria failed: %v", result.Error)
	} else if (result.RowsAffected == 0) {
		t.Errorf("putFerreteria had 0 rows")
	}

	err := mock.ExpectationsWereMet()
	if (err != nil) {
		t.Errorf("Failed to meet expectations, got error: %v", err)
	}

	teardown(gdb)
}
