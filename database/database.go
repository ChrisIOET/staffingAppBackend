package database

import (
	"fmt"
	"projectStaff/internal/envvar"
	"projectStaff/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

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

func MigrateDatabaseTables() {

	DB.AutoMigrate(&model.Skill{})
	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.UserSkill{})
	DB.AutoMigrate(&model.Category{})
}






