package main

import (
	"context"
	"log"

	controller "GoProject/controller"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoURI   = "mongodb://localhost:27017"
	dbName     = "project"
	collection = "users"
)

func main() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	db := client.Database(dbName)
	usersCollection := db.Collection(collection)

	router := gin.Default()

	userController := controller.SetupUserController(usersCollection)

	// Update routes to include :id parameter
	router.POST("/users", userController.CreateUserHandler)
	router.GET("/users", userController.ReadUserHandler)
	router.PUT("/users/:id", userController.UpdateUserHandler)
	router.DELETE("/users/:id", userController.DeleteUserHandler)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
