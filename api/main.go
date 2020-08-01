package main

import (
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
	return router.SetupRoute(taskController)
}

func main() {
	setupServer("dev").Run(":3000")
}
