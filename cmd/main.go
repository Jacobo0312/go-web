package main

import (
	"database/sql"
	"log"

	"github.com/Jacobo0312/go-web/cmd/server"
	"github.com/Jacobo0312/go-web/config"
	"github.com/Jacobo0312/go-web/pkg/firebase"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Load error config: %v", err)
	}

	//Firebase connection
	firebase.InitFirebase()

	// DB connection
	log.Println("Connecting to database...")
	db, err := sql.Open("mysql", cfg.DBConnString)
	if err != nil {
		log.Fatalf("Error open database: %v", err)
	}
	defer db.Close()

	// Run migrations
	log.Println("Running migrations...")
	if err := runMigrations(db); err != nil {
		log.Fatalf("Error running migrations: %v", err)
	}

	srv := server.New(cfg, db)

	if err := srv.Start(); err != nil {
		log.Fatalf("Starting server error: %v", err)
	}
}

func runMigrations(db *sql.DB) error {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"mysql", driver)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
