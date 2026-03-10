package database

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/golang-migrate/migrate/v4"
	migratepostgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
    host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")


    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
        host, user, password, dbname, port, sslmode,
    )
    // dsn := "host=localhost user=orderly-admin password=admin123 dbname=orderly_auth port=5432 sslmode=disable"

    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // Run migrations
    sqlDB, _ := DB.DB()
    driver, _ := migratepostgres.WithInstance(sqlDB, &migratepostgres.Config{})

	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	migrationPath := filepath.ToSlash(filepath.Join(basepath, "migrations"))
	sourceURL := "file://" + migrationPath
    
    m, err := migrate.NewWithDatabaseInstance(
        sourceURL, 
        "postgres", 
        driver,
    )
    if err != nil {
        log.Fatal("Migration setup failed:", err)
    }

    if err := m.Up(); err != nil && err != migrate.ErrNoChange {
        log.Fatal("Migration failed:", err)
    }
}