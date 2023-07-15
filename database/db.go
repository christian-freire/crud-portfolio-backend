package database

import (
	"log"
	"projetos-pessoais/crud-porfolio-backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectToDataBase() {
	connectionString := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		log.Panic("Error to connect to databases")
	}
	DB.AutoMigrate(&models.User{})
}
