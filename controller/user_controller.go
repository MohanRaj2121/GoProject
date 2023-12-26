package controller

import (
	model "GoProject/model"
	service "GoProject/service"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct {
	Service *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{Service: service}
}

// create user
func (c *UserController) CreateUserHandler(ctx *gin.Context) {
	var user model.UserRequest
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go c.Service.CreateUser(user, &wg)
	wg.Wait()

	ctx.JSON(http.StatusCreated, gin.H{"message": "user created"})
}

// read
func (c *UserController) ReadUserHandler(ctx *gin.Context) {
	var wg sync.WaitGroup
	resultCh := make(chan []model.User)

	wg.Add(1)
	go c.Service.ReadUsers(&wg, resultCh)

	wg.Wait()
	close(resultCh)
	go func() {

	}()

	users := <-resultCh
	ctx.JSON(http.StatusOK, users)
}

// update

func (c *UserController) UpdateUserHandler(ctx *gin.Context) {
	userID := ctx.Param("id")
	var updateData struct {
		NewEmail string `json:"new_email"`
	}
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go c.Service.UpdateUser(userID, updateData.NewEmail, &wg)
	wg.Wait()

	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// delete

func (c *UserController) DeleteUserHandler(ctx *gin.Context) {
	userID := ctx.Param("id")

	var wg sync.WaitGroup
	wg.Add(1)
	go c.Service.DeleteUser(userID, &wg)
	wg.Wait()

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func SetupUserController(collection *mongo.Collection) *UserController {
	userService := service.NewUserService(collection)
	return NewUserController(userService)
}
