package config

import (
    "database/sql"
    "fmt"
    "os"
    "time"
)

var (
    databaseInstance *sql.DB
    errorInit        error
)

func PostgresConnect() (*sql.DB, error) {

    if databaseInstance != nil {
        return databaseInstance, nil
    }

    host := os.Getenv("POSTGRES_HOST")
    port := os.Getenv("POSTGRES_PORT")
    user := os.Getenv("POSTGRES_USER")
    password := os.Getenv("POSTGRES_PASSWORD")
    dbname := os.Getenv("POSTGRES_DB")

    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

    database, err := sql.Open("postgres", dsn)

    if err != nil {
        errorInit = fmt.Errorf("error opening database: %v", err)
        return nil, errorInit
    }

    database.SetMaxOpenConns(25)
    database.SetMaxIdleConns(25)
    database.SetConnMaxLifetime(5 * time.Minute)

    if err = database.Ping(); err != nil {
        return nil, fmt.Errorf("error pinging database: %v", err)
    }

    fmt.Println("Postgres connected")

    databaseInstance = database

    return databaseInstance, errorInit
}
