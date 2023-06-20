package models

type Student struct {
	ID             int    `db:"id" json:"id"`
	FirstName      string `db:"first_name" json:"first_name"`
	LastName       string `db:"last_name" json:"last_name"`
	Email          string `db:"email" json:"email"`
	EnrollmentYear int    `db:"enrollment_year" json:"enrollment_year"`
}
