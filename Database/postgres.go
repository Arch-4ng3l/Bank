package database

import (
	"database/sql"
	"fmt"
	"os"
)

type Postgres struct {
	Db     *sql.DB
	dbName string
	host   string
	port   uint16
}

func NewPostgres(port uint16, host string, dbName string) *Postgres {
	return &Postgres{
		dbName: dbName,
		host:   host,
		port:   port,
	}
}

func (psql *Postgres) Connect() error {
	password := os.Getenv("PSQLPW")
	username := os.Getenv("PSQLUN")

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		psql.host, psql.port, username, password, psql.dbName)

	if db, err := sql.Open("postgres", connStr); err != nil {
		return err
	} else {
		if err := db.Ping(); err != nil {
			psql.Db = db
		} else {
			return err
		}
	}

	return nil
}
