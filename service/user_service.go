// service/user_service.go

package service

import (
	model "GoProject/model"
	"context"
	"fmt"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	Collection *mongo.Collection
}

func NewUserService(collection *mongo.Collection) *UserService {
	return &UserService{Collection: collection}
}

func (s *UserService) CreateUser(user model.UserRequest, wg *sync.WaitGroup) {
	defer wg.Done()
	userModel := model.User{
		UserName: user.UserName,
		ID:       user.UserName,
		Email:    user.UserName,
	}
	_, err := s.Collection.InsertOne(context.Background(), userModel)
	if err != nil {
		log.Println("Error creating user:", err)
	}
	fmt.Println("User created successfully")
}

func (s *UserService) ReadUsers(wg *sync.WaitGroup, resultCh chan<- []model.User) {
	defer wg.Done()

	var users []model.User
	cursor, err := s.Collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Println("Error reading users:", err)
		return
	}
	defer cursor.Close(context.Background())

	err = cursor.All(context.Background(), &users)
	if err != nil {
		log.Println("Error reading users:", err)
		return
	}
	resultCh <- users
}

func (s *UserService) UpdateUser(userID, newEmail string, wg *sync.WaitGroup) {
	defer wg.Done()

	filter := bson.D{{Key: "_id", Value: userID}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "email", Value: newEmail}}}}

	_, err := s.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println("Error updating user:", err)
		return
	}
	fmt.Println("User updated successfully")
}

func (s *UserService) DeleteUser(userID string, wg *sync.WaitGroup) {
	defer wg.Done()

	filter := bson.D{{Key: "_id", Value: userID}}

	_, err := s.Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Println("Error deleting user:", err)
		return
	}
	fmt.Println("User deleted successfully")
}
