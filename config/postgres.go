package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"sync"
	"time"
)

var (
	databaseInstance *sql.DB
	oncePostgres     sync.Once
	errorInit          error
)

func PostgresConnect() (*sql.DB, error) {
	oncePostgres.Do(func() {
		host := os.Getenv("POSTGRES_HOST")
		port := os.Getenv("POSTGRES_PORT")
		user := os.Getenv("POSTGRES_USER")
		password := os.Getenv("POSTGRES_PASSWORD")
		dbname := os.Getenv("POSTGRES_DB")

		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

		db, err := sql.Open("postgres", dsn)
		if err != nil {
			errorInit = fmt.Errorf("error opening database: %v", err)
			return
		}

		db.SetMaxOpenConns(25)
		db.SetMaxIdleConns(25)
		db.SetConnMaxLifetime(5 * time.Minute)

		if err = db.Ping(); err != nil {
			errorInit = fmt.Errorf("error pinging database: %v", err)
			return
		}

		fmt.Println("Connected to Postgres")
		databaseInstance = db
	})

	return databaseInstance, errorInit
}
