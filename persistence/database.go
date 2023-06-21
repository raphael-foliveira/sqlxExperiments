package persistence

import (
	"database/sql"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

type DbCredentials struct {
	Host     string
	User     string
	Password string
	DbName   string
}

func InitDb(c *DbCredentials) (*sqlx.Tx, error) {
	db, err := connectDb(c)
	if err != nil {
		return nil, err
	}
	_, err = migrateSchema(db)
	if err != nil {
		return nil, err
	}
	return db.Beginx()
}

func connectDb(c *DbCredentials) (*sqlx.DB, error) {
	return sqlx.Connect(
		"postgres",
		"host="+c.Host+
			" user="+c.User+
			" password="+c.Password+
			" dbname="+c.DbName+
			" sslmode=disable")
}

func migrateSchema(db *sqlx.DB) (sql.Result, error) {
	return db.Exec(Schema)
}
