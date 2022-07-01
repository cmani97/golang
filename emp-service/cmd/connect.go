package cmd

import (
	"fmt"

	"github.com/cmani097/emp-service/models/entity"
	// "gorm.io/driver/postgres"
	// "gorm.io/gorm"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func ConnectToPostgres() *gorm.DB {
	// dbURL := "postgres://postgres:postgresql@localhost:5432/slice"
	// db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	db, err := gorm.Open("postgres", "user=postgres password=postgresql dbname=slice sslmode=disable")

	if err != nil {
		fmt.Println("err= ", err)
		panic("Failed to connect to database!")
	}

	// db.DropTableIfExists(&entity.Employee{}, &entity.Team{}, &entity.Address{})
	db.AutoMigrate(&entity.Employee{})
	db.AutoMigrate(&entity.Team{})
	db.AutoMigrate(&entity.Address{})
	return db
}
