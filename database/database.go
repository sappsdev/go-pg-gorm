package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	pg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {

	db_user := os.Getenv("DB_USERNAME")
	db_pass := os.Getenv("DB_PASSWORD")
	db_host := os.Getenv("DB_HOST")
	db_name := os.Getenv("DB_NAME")
	db_port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=disable", db_host, db_user, db_pass, db_name, db_port)
	runMigrations(dsn)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func runMigrations(dsn string) {
	db, _ := sql.Open("postgres", dsn)
	driver, _ := pg.WithInstance(db, &pg.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://database/migrations",
		"postgres", driver)
	m.Up()
}
