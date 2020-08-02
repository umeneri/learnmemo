package router

import (
	"api/interfaces/controller"
	"api/interfaces/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoute(taskController controller.TaskController, userController controller.UserController) *gin.Engine {
	engine := gin.Default()
	engine.Use(middleware.RecordUaAndTime)
	engine.LoadHTMLGlob("web/*")
	taskRoute := engine.Group("/task")
	{
		v1 := taskRoute.Group("/v1")
		{
			v1.POST("/add", taskController.AddTask)
			v1.GET("/list", taskController.ListTask)
			v1.PUT("/update/:id", taskController.UpdateTask)
			v1.DELETE("/delete/:id", taskController.DeleteTask)
		}
	}
	userRoute := engine.Group("/user")
	{
		v1 := userRoute.Group("/v1")
		{
			v1.GET("/index", userController.Index)
			v1.GET("/login", userController.Login)
			v1.GET("/google-callback", userController.GoogleCallback)
		}
	}
	return engine

}
