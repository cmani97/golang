package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/postgres" // The purpose of _ is for import for side effect, without any explicit use.
	"github.com/spf13/viper"
)

func SetupModels() *gorm.DB {
	//db, err := gorm.Open("sqlite3", "test.db")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	// no need - start
	viper_user := viper.Get("POSTGRES_USER")
	viper_password := viper.Get("POSTGRES_PASSWORD")
	viper_db := viper.Get("POSTGRES_DB")
	viper_host := viper.Get("POSTGRES_HOST")
	viper_port := viper.Get("POSTGRES_PORT")
	fmt.Println("viper_user= ", viper_user)
	fmt.Println("viper_password= ", viper_password)
	fmt.Println("viper_db= ", viper_db)
	fmt.Println("viper_host= ", viper_host)
	fmt.Println("viper_port= ", viper_port)
	// no need - end

	dbURL := "postgres://postgres:postgresql@localhost:5432/books_store"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		fmt.Println("err= ", err)
		panic("Failed to connect to database!")
	}
	db.AutoMigrate(&Book{})

	return db
}
