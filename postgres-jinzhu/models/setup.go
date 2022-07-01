package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // The purpose of _ is for import for side effect, without any explicit use.
	"github.com/spf13/viper"
)

// Customer ...
type Customer struct {
	CustID       int `gorm:"primary_key"`
	CustomerName string
	Contacts     []Contact `gorm:"foreignKey:CustID;references:CustID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// Contact ...
type Contact struct {
	ContactID   int `gorm:"primary_key"`
	CountryCode int
	MobileNo    uint `gorm:"size:13;not null"`
	CustID      int
	AuthID      uint
}

func SetupModels() *gorm.DB {
	//db, err := gorm.Open("sqlite3", "test.db")

	viper.SetConfigFile(".env")
	viper.AddConfigPath(".")
	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	// no need - start
	viper_user := viper.GetString("POSTGRES_USER")
	viper_password := viper.GetString("POSTGRES_PASSWORD")
	viper_db := viper.GetString("POSTGRES_DB")
	viper_host := viper.GetString("POSTGRES_HOST")
	viper_port := viper.Get("POSTGRES_PORT")
	fmt.Println("viper_user= ", viper_user)
	fmt.Println("viper_password= ", viper_password)
	fmt.Println("viper_db= ", viper_db)
	fmt.Println("viper_host= ", viper_host)
	fmt.Println("viper_port= ", viper_port)
	// no need - end

	// dbURL := "postgres://postgres:postgresql@localhost:5432/books_store"
	// db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	db, err := gorm.Open("postgres", "user=postgres password=postgresql dbname=books_store sslmode=disable")
	if err != nil {
		fmt.Println("err= ", err)
		panic("Failed to connect to database!")
	}

	defer db.Close()
	db.DropTableIfExists(&Book{}, &Author{})
	db.AutoMigrate(&Book{}, &Author{})
	db.Model(&Book{}).Joins("authors on authors.id = books.author_id")
	// db.Model(&Book{}).Related(&Author{})
	// db.Model(&Contact{}).AddForeignKey("auth_id", "authors(auth_id)", "CASCADE", "CASCADE")

	db.DropTableIfExists(&Contact{}, &Customer{})
	db.AutoMigrate(&Customer{}, &Contact{})
	// db.Model(&Contact{}).AddForeignKey("cust_id", "customers(cust_id)", "CASCADE", "CASCADE")

	Custs1 := Customer{CustomerName: "John", Contacts: []Contact{
		{CountryCode: 91, MobileNo: 956112},
		{CountryCode: 91, MobileNo: 997555}}}

	Custs2 := Customer{CustomerName: "Martin", Contacts: []Contact{
		{CountryCode: 90, MobileNo: 808988},
		{CountryCode: 90, MobileNo: 909699}}}
	db.Create(&Custs1)
	db.Create(&Custs2)

	book := Book{Title: "Clean Code", Desc: "A Handbook of Agile Software Craftsmanship",
		Author: Author{Name: "Robert C. Martin Series", Degree: "Ph.D", Contacts: []Contact{
			{CountryCode: 91, MobileNo: 956112},
			{CountryCode: 91, MobileNo: 997555}}}}

	book2 := Book{Title: "Clean Code1", Desc: "A Handbook of Agile Software Craftsmanship1",
		Author: Author{Name: "Robert C. Martin Series1", Degree: "Ph.D1", Contacts: []Contact{
			{CountryCode: 91, MobileNo: 9561121},
			{CountryCode: 91, MobileNo: 9975551}}}}

	db.Save(&book)
	db.Save(&book2)
	cust := []Customer{}
	db.Debug().Preload("Contacts").Find(&cust)
	fmt.Println("cust= ", cust)

	var book1 Book
	var authors Author
	// db.Preload("Authors").Find(&allBooks)
	db.Model(&book1).Association("Authors").Find(&authors)
	fmt.Println("allBooks= ", book1)

	// db.Preload("Contacts").Find(&authors)
	// query := db.Table("authors").Select("*").Joins("left join contacts contacts on contacts.auth_id = authors.id")
	// fmt.Println("query= ", query)
	// b := []Book{}
	// db.Model(&book1).Joins("join (?) q on q.id = books.author_id", query).Scan(&b)
	// fmt.Println("books b= ", b)

	return db
}
