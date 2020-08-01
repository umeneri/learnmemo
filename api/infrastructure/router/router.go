package router

import (
	"api/interfaces/controller"
	"api/interfaces/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoute(taskController controller.TaskController) *gin.Engine {
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

