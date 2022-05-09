package controllers

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ferromarket/backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

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

	mock.ExpectExec(
		regexp.QuoteMeta("INSERT INTO `ferreteria` (`name`) VALUES (?)")).
		WithArgs("Chris's hardware").
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	postFerreteria(models.Ferreteria{Name: "Chris's hardware"}, db)

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Failed to meet expectations, got error: %v", err)
	}
}
