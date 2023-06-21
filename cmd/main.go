package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/raphael-foliveira/sqlxExperiments/persistence"
	"github.com/raphael-foliveira/sqlxExperiments/persistence/models"
	"github.com/raphael-foliveira/sqlxExperiments/persistence/repositories"
)

var dbCredentials = persistence.DbCredentials{
	Host:     "localhost",
	User:     "postgres",
	Password: "postgres",
	DbName:   "sqlxtest",
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}
	run()
}

func run() *models.Student {
	tx, err := persistence.InitDb(&dbCredentials)
	if err != nil {
		panic(err.Error())
	}
	studentRepository := repositories.NewStudentRepository(tx)

	johnDoe := models.Student{FirstName: "John", LastName: "Doe", Email: "john@doe.com", EnrollmentYear: 2019}

	err = studentRepository.Create(&johnDoe)
	if err != nil {
		fmt.Println(err.Error(), "could not add to database")
	}

	student, err := studentRepository.GetByFirstName("John")

	fmt.Println(student)
	return student
}
