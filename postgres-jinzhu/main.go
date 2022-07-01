package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/slice-mani/postgres-jinzhu/controllers"
	"github.com/slice-mani/postgres-jinzhu/models"
	"github.com/slice-mani/postgres-jinzhu/services"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server           *gin.Engine
	database         *gorm.DB
	err              error
	ctx              context.Context
	usercollection   *mongo.Collection
	mongoclient      *mongo.Client
	bookcontroller   controllers.BookController
	samplecontroller controllers.SampleController
	sampleservice    services.SampleService
	bookservice      services.BookService
)

func init() {
	server = gin.Default()
	database = models.SetupModels()

	//test
	mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal(err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mongo connection established")

	usercollection = (*mongo.Collection)(mongoclient.Database("userdb").Collection("users"))
	// end

	bookservice = services.NewBookService(database, ctx)
	bookcontroller = controllers.NewBookController(bookservice)
	sampleservice = services.NewSampleService(usercollection, ctx)
	samplecontroller = controllers.New(sampleservice)
}
func main() {

	// Provide db variable to controllers
	server.Use(func(c *gin.Context) {
		c.Set("db", database)
		c.Next()
	})

	server.GET("/test", samplecontroller.TestService)
	server.GET("/books", bookcontroller.FindBooks)
	server.POST("/books", bookcontroller.CreateBook)
	server.GET("/books/:id", bookcontroller.FindBook)
	server.PATCH("/books/:id", bookcontroller.UpdateBook)
	server.DELETE("/books/:id", bookcontroller.DeleteBook)
	server.Run(":8085")
}
