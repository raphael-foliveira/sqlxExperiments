package repositories

import (
	"log"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/raphael-foliveira/sqlxExperiments/persistence/models"
)

var db *sqlx.DB
var studentId int64

func TestMain(m *testing.M) {
	var err error
	db, err = sqlx.Connect("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	tx := db.MustBegin()
	_, err = tx.Exec(`
		CREATE TABLE IF NOT EXISTS student (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		first_name TEXT,
		last_name TEXT,
		email TEXT,
		enrollment_year INTEGER
	);
	INSERT INTO student (first_name, last_name, email, enrollment_year) VALUES ('Spam', 'Eggs', 'spam@eggs.com', 2019);`)
	if err != nil {
		log.Fatal(err, "could not create table")
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err, "could not commit")
	}
	m.Run()
	defer db.Close()
}

func TestCreate(t *testing.T) {

	studentRepository := NewStudentRepository(db.MustBegin())
	student := &models.Student{
		FirstName:      "John",
		LastName:       "Doe",
		Email:          "john@doe.com",
		EnrollmentYear: 2019,
	}

	err := studentRepository.Create(student)
	if err != nil {
		t.Error("could not add to database")
	}
}

func TestGetById(t *testing.T) {
	tx := db.MustBegin()
	sr := NewStudentRepository(tx)
	student, err := sr.GetById(1)
	tx.Commit()
	if err != nil {
		t.Errorf("could not get student %d from database", 1)
	}
	if student.FirstName != "Spam" {
		t.Error("wrong student")
	}
}

func TestGetByFirstName(t *testing.T) {
	tx := db.MustBegin()
	sr := NewStudentRepository(tx)
	student, err := sr.GetByFirstName("Spam")
	tx.Commit()
	if err != nil {
		t.Errorf("could not get student %s from database", "Spam")
	}
	if student.FirstName != "Spam" {
		t.Error("wrong student")
	}
}
