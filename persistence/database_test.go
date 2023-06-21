package persistence

import "testing"

var dbCredentials = &DbCredentials{
	Host:     "localhost",
	User:     "postgres",
	Password: "postgres",
	DbName:   "sqlxtest",
}

func TestInitDb(t *testing.T) {
	_, err := InitDb(dbCredentials)
	if err != nil {
		t.Error("could not connect to database")
	}
}
