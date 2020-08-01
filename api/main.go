package main

import (
	_ "github.com/go-sql-driver/mysql"

	"api/usecase"
	"api/interfaces/controller"
	"api/infrastructure/router"
	"api/infrastructure/repository"
)

func setupServer() {
	dbName := "gin"
	taskRepository := repository.NewTaskRepository(dbName)
	taskUseCase := usecase.NewTaskUseCase(taskRepository)
	taskController := controller.NewTaskController(taskUseCase)
	router.SetupRoute(taskController).Run(":3000")
}

func main() {
	setupServer()
}
