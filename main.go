package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"api/infrastructure/repository"
	"api/infrastructure/router"
	"api/interfaces/controller"
	"api/usecase"
)

func setupServer() *gin.Engine {
	taskRepository := repository.NewTaskRepository()
	taskUseCase := usecase.NewTaskUseCase(taskRepository)
	taskController := controller.NewTaskController(taskUseCase)
	userRepository := repository.NewUserRepository()
	userUseCase := usecase.NewUserUseCase(userRepository)
	userController := controller.NewUserController(userUseCase)
	return router.SetupRoute(taskController, userController)
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	setupServer().Run(fmt.Sprintf(":%s", port))
}
