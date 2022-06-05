package utils

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"strings"
	"time"
)

type Date time.Time // 2006-01-02
type DateTime time.Time // 2006-01-02 15:04:05

func (j *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if (err != nil) {
		return err
	}
	*j = Date(t)
	return nil
}

func (j Date) MarshalJSON() ([]byte, error) {
    return json.Marshal(time.Time(j))
}

func (j Date) Format(s string) string {
    t := time.Time(j)
    return t.Format(s)
}

func (date *Date) Scan(value interface{}) (err error) {
	nullTime := &sql.NullTime{}
	err = nullTime.Scan(value)
	*date = Date(nullTime.Time)
	return
}

func (date Date) Value() (driver.Value, error) {
	return time.Time(date), nil
}

func (date Date) GormDataType() string {
	return "date"
}

func (date Date) GobEncode() ([]byte, error) {
	return time.Time(date).GobEncode()
}

func (date *Date) GobDecode(b []byte) error {
	return (*time.Time)(date).GobDecode(b)
}

func (j *DateTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02 15:04:05", s)
	if (err != nil) {
		return err
	}
	*j = DateTime(t)
	return nil
}

func (j DateTime) MarshalJSON() ([]byte, error) {
    return json.Marshal(time.Time(j))
}

func (j DateTime) Format(s string) string {
    t := time.Time(j)
    return t.Format(s)
}

func (date *DateTime) Scan(value interface{}) (err error) {
	nullTime := &sql.NullTime{}
	err = nullTime.Scan(value)
	*date = DateTime(nullTime.Time)
	return
}

func (date DateTime) Value() (driver.Value, error) {
	return time.Time(date), nil
}

func (date DateTime) GormDataType() string {
	return "datetime"
}

func (date DateTime) GobEncode() ([]byte, error) {
	return time.Time(date).GobEncode()
}

func (date *DateTime) GobDecode(b []byte) error {
	return (*time.Time)(date).GobDecode(b)
}
