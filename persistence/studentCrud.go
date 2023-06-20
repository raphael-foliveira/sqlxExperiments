package persistence

import (
	"github.com/jmoiron/sqlx"
	"github.com/raphael-foliveira/sqlxExperiments/models"
)

type StudentContextManager struct {
	Tx *sqlx.Tx
}

func NewStudentContextManager(tx *sqlx.Tx) *StudentContextManager {
	return &StudentContextManager{Tx: tx}
}

func (scm *StudentContextManager) Create(s *models.Student) error {
	_, err := scm.Tx.Exec(
		`INSERT INTO student (first_name, last_name, email, enrollment_year)
		VALUES ($1, $2, $3, $4)`,
		s.FirstName,
		s.LastName,
		s.Email,
		s.EnrollmentYear)
	if err != nil {
		return err
	}
	return scm.Tx.Commit()
}

func (scm *StudentContextManager) GetById(id int) (*models.Student, error) {
	var student models.Student
	err := scm.Tx.Get(&student, "SELECT * FROM student WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (scm *StudentContextManager) GetByFirstName(firstName string) (*models.Student, error) {
	var student models.Student
	err := scm.Tx.Get(&student, "SELECT * FROM student WHERE first_name=$1", firstName)
	if err != nil {
		return nil, err
	}
	return &student, nil
}
