package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"api/controller"
	"api/middleware"
	"api/service"
)

func setupServer() *gin.Engine {
	taskService := service.NewTaskService()
	taskController := controller.NewTaskController(taskService)

	engine := gin.Default()
	engine.Use(middleware.RecordUaAndTime)
	taskEngine := engine.Group("/task")
	{
		v1 := taskEngine.Group("/v1")
		{
			v1.POST("/add", taskController.AddTask)
			v1.GET("/list", taskController.ListTask)
			v1.PUT("/update/:id", taskController.UpdateTask)
			v1.DELETE("/delete/:id", taskController.DeleteTask)
		}
	}
	return engine
}

func main() {
	setupServer().Run(":3000")
}
