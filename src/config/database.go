package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func PostgresConnect() (*sql.DB, error) {
	host := postgresHost
	port := postgresPort
	user := postgresUser
	pass := postgresPass
	dbname := postgresDbName

	conString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)
	db, err := sql.Open("postgres", conString)

	if err != nil {
		return db, err
	}
	return db, nil
}
