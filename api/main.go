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

func setupServer(env string) *gin.Engine {
	dbName := "gin"
	if env == "test" {
		dbName = "gin_test"
	}
	taskRepository := repository.NewTaskRepository(dbName)
	taskUseCase := usecase.NewTaskUseCase(taskRepository)
	taskController := controller.NewTaskController(taskUseCase)
	userRepository := repository.NewUserRepository(dbName)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userController := controller.NewUserController(userUseCase)
	return router.SetupRoute(taskController, userController)
}

func main() {
	env := os.Getenv("ENV")
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	setupServer(env).Run(fmt.Sprintf(":%s", port))
}
