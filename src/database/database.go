package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {
	dsn := "host=localhost user=postgres password=123456 dbname=db_usuarios port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Erro ao conectar à DataBase")
	}
	db = database
}

func GetDatabase() *gorm.DB {
	return db
}
