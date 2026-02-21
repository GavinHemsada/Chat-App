package database

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func RunMigrations(dsn string) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Dynamic migration path finding
	migrationsPath := "migrations" // default
	searchPaths := []string{
		"migrations",
		"../migrations",
		"../../migrations",
		"../../../migrations",
	}

	for _, path := range searchPaths {
		if _, err := os.Stat(path); err == nil {
			migrationsPath = path
			break
		}
	}

	log.Printf("Using migrations from: %s", migrationsPath)

	pathToUse := filepath.ToSlash(migrationsPath)
	
	m, err := migrate.NewWithDatabaseInstance(
		"file://" + pathToUse,
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	log.Println("Migrations applied successfully")
}
