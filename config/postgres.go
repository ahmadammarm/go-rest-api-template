package config

import (
	"database/sql"
	"fmt"
	"os"
	"sync"
	"time"
)

var (
	databaseInstance *sql.DB
	postgresOnce     sync.Once
	errorInit        error
)

func PostgresConnect() (*sql.DB, error) {
	postgresOnce.Do(func() {
		host := os.Getenv("POSTGRES_HOST")
		port := os.Getenv("POSTGRES_PORT")
		user := os.Getenv("POSTGRES_USER")
		password := os.Getenv("POSTGRES_PASSWORD")
		dbname := os.Getenv("POSTGRES_DB")

		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

		database, err := sql.Open("postgres", dsn)

		if err != nil {
			errorInit = fmt.Errorf("error connecting to database: %v", err)
			return
		}

		database.SetMaxOpenConns(25)

		database.SetMaxIdleConns(25)

		database.SetConnMaxLifetime(5 * time.Minute)

		if err = database.Ping(); err != nil {
			errorInit = fmt.Errorf("error pinging database: %v", err)
			return
		}

		fmt.Println("Postgres connected")

		databaseInstance = database

	})

	return databaseInstance, errorInit
}
