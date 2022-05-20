package controllers

import (
	"database/sql/driver"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ferromarket/backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func TestPostFerreteria(t *testing.T) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	if sqlDB == nil {
		t.Error("mock db is null")
	}
	if mock == nil {
		t.Error("sqlmock is null")
	}
	defer sqlDB.Close()

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "sqlmock_db_0",
		DriverName: "mysql",
		Conn: sqlDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		t.Errorf("Failed to open gorm v2 db, got error: %v", err)
	}

	mock.ExpectBegin()

	ferreteria := models.Ferreteria{
		Nombre: "Chris's hardware",
		ComunaID: 1,
		Direccion: "214 Edison Dr.",
		Descripcion: "blah blah blah",
	}

	mock.ExpectExec(
		regexp.QuoteMeta("INSERT INTO `ferreteria` (`created_at`,`updated_at`,`deleted_at`,`nombre`,`comuna_id`,`direccion`,`descripcion`) VALUES (?,?,?,?,?,?,?)")).
		WithArgs(AnyTime{}, AnyTime{}, nil, ferreteria.Nombre, ferreteria.ComunaID, ferreteria.Direccion, ferreteria.Descripcion).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	postFerreteria(ferreteria, db)

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Failed to meet expectations, got error: %v", err)
	}
}
