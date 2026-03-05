package database

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	migratepostgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

    dsn := "host=localhost user=orderly-admin password=admin123 dbname=orderly_oms port=5432 sslmode=disable"

    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // Run migrations
    sqlDB, _ := DB.DB()
    driver, _ := migratepostgres.WithInstance(sqlDB, &migratepostgres.Config{})
    
    m, err := migrate.NewWithDatabaseInstance(
        "file://../../migrations",
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