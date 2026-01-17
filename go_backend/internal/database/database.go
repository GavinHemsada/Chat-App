package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

// Connect opens a database connection and returns *pgx.Conn
func Connect(host, user, password, dbname string) *pgx.Conn {
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s", user, password, host, dbname)

	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connected!")
	return conn
}