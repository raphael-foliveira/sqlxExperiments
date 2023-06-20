package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/raphael-foliveira/sqlxExperiments/models"
	"github.com/raphael-foliveira/sqlxExperiments/persistence"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	tx, err := persistence.InitDb()
	if err != nil {
		panic(err.Error())
	}
	studentCm := persistence.NewStudentContextManager(tx)

	johnDoe := models.Student{FirstName: "John", LastName: "Doe", Email: "john@doe.com", EnrollmentYear: 2019}

	err = studentCm.Create(&johnDoe)
	if err != nil {
		fmt.Println(err.Error(), "could not add to database")
	}

	student, err := studentCm.GetByFirstName("John")

	fmt.Println(student)
}
