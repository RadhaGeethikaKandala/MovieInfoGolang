package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)


func CreateDatabaseConn() *sql.DB{
	dataSource := fmt.Sprintf(
		"postgres://%s:%s@localhost:5432/%s?sslmode=disable",
		"movierental",
		"movierental",
		"movie_rental",
	)

	dbConn, err := sql.Open("postgres", dataSource)
	errString := "Unable to open a connection to database, error: %s"
	if err != nil {
		log.Fatalf(errString, err.Error())
	}
	if err := dbConn.Ping(); err != nil {
		log.Fatalf(errString, err.Error())
	}

	return dbConn
}

func RunDatabaseMigrations() {
	dataSource := fmt.Sprintf(
		"postgres://%s:%s@localhost:5432/%s?sslmode=disable",
		"movierental",
		"movierental",
		"movie_rental",
	)

	m, err := migrate.New("file://internal/app/database/migrations", dataSource)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}


}