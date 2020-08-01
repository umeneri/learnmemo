package main

import (
	_ "github.com/go-sql-driver/mysql"

	"api/interfaces/controller"
	"api/interfaces/router"
	"api/service"
)

func setupServer() {
	dbName := "gin"
	taskService := service.NewTaskService(dbName)
	taskController := controller.NewTaskController(taskService)
	router.SetupRoute(taskController).Run(":3000")
}

func main() {
	setupServer()
}
