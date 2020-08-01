package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"api/controller"
	"api/middleware"
	"api/service"
)

func setupRoute(taskController controller.TaskController) *gin.Engine {
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
func setupServer() *gin.Engine {
	dbName := "gin"
	taskService := service.NewTaskService(dbName)
	taskController := controller.NewTaskController(taskService)

	return setupRoute(taskController)
}

func main() {
	setupServer().Run(":3000")
}
