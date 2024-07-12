package postgres

import (
	"database/sql"
	"fmt"
	"log"
)

var ErrNoMatch = fmt.Errorf("no matching record")

type Database struct {
	Conn *sql.DB
}

type DBParams struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func Initialize(params DBParams) (*Database, error) {
	db := &Database{}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		params.Host, params.Port, params.Username, params.Password, params.Database)
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return db, err
	}

	db.Conn = conn
	err = db.Conn.Ping()
	if err != nil {
		return db, err
	}

	log.Println("Database connection established")

	return db, nil
}
