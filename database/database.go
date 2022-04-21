package database

import (
	"fmt"
	"projectStaff/internal/envvar"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// GetDatabase returns a database instance.
func GetDatabase() *gorm.DB {
	
	user := envvar.DBUser()
	password := envvar.DBPassword()
	dbname := envvar.DBName()
	dbhost := envvar.DBHost()
	dbport := envvar.DBPort()

	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		user, password, dbname, dbhost, dbport)
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db
}
