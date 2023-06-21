package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/raphael-foliveira/sqlxExperiments/persistence/models"
)

type StudentRepository struct {
	Tx *sqlx.Tx
}

func NewStudentRepository(tx *sqlx.Tx) *StudentRepository {
	return &StudentRepository{Tx: tx}
}

func (sr *StudentRepository) Create(s *models.Student) error {
	_, err := sr.Tx.Exec(
		`INSERT INTO student (first_name, last_name, email, enrollment_year)
		VALUES ($1, $2, $3, $4)`,
		s.FirstName,
		s.LastName,
		s.Email,
		s.EnrollmentYear)
	if err != nil {
		return err
	}
	return sr.Tx.Commit()
}

func (sr *StudentRepository) GetById(id int) (*models.Student, error) {
	var student models.Student
	err := sr.Tx.Get(&student, "SELECT first_name, last_name, email, enrollment_year FROM student WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (sr *StudentRepository) GetByFirstName(firstName string) (*models.Student, error) {
	var student models.Student
	err := sr.Tx.Get(&student, "SELECT first_name, last_name, email, enrollment_year FROM student WHERE first_name=$1", firstName)
	if err != nil {
		return nil, err
	}
	return &student, nil
}
