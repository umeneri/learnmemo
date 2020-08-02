package router

import (
	"api/interfaces/controller"
	// "api/interfaces/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoute(taskController controller.TaskController, userController controller.UserController) *gin.Engine {
	engine := gin.Default()
	// engine.Use(middleware.RecordUaAndTime)
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

	engine.GET("/", userController.Index)

	userRoute := engine.Group("/user")
	{
		v1 := userRoute.Group("/v1")
		{
			v1.GET("/login", userController.LoginIndex)
			v1.GET("/auth/:provider", userController.Login)
			v1.GET("/callback/:provider", userController.Callback)
		}
	}
	return engine

}
