package config

import (
	"database/sql"
	"fmt"
	"os"
	"time"
)

func PostgresConnect() (*sql.DB, error) {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		return nil, fmt.Errorf("one or more required environment variables are missing")
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	databaseInstance, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	databaseInstance.SetMaxOpenConns(25)
	databaseInstance.SetMaxIdleConns(25)
	databaseInstance.SetConnMaxLifetime(5 * time.Minute)

	if err = databaseInstance.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %v", err)
	}

	fmt.Println("Postgres connected")
	return databaseInstance, nil
}
