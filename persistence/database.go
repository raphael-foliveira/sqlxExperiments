package persistence

import (
	"database/sql"
	"os"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

func InitDb() (*sqlx.Tx, error) {
	db, err := connectDb()
	if err != nil {
		return nil, err
	}
	_, err = migrateSchema(db)
	if err != nil {
		return nil, err
	}
	return db.Beginx()

}

func connectDb() (*sqlx.DB, error) {
	return sqlx.Connect(
		"postgres",
		"host=localhost"+
			" user="+os.Getenv("POSTGRES_USER")+
			" password="+os.Getenv("POSTGRES_PASSWORD")+
			" dbname="+os.Getenv("POSTGRES_DB")+
			" sslmode=disable")
}

func migrateSchema(db *sqlx.DB) (sql.Result, error) {
	return db.Exec(schema)
}
