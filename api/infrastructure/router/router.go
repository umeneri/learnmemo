package router

import (
	"api/interfaces/auth"
	"api/interfaces/controller"

	"api/interfaces/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoute(taskController controller.TaskController, userController controller.UserController) *gin.Engine {
	engine := gin.Default()
	engine.Use(middleware.RecordUaAndTime)
	engine.LoadHTMLGlob("web/*")
	taskRoute := engine.Group("/api/task")
	{
		v1 := taskRoute.Group("/v1")
		{
			v1.POST("/add", taskController.AddTask)
			v1.GET("/list", taskController.ListTask)
			v1.PUT("/update/:id", taskController.UpdateTask)
			v1.DELETE("/delete/:id", taskController.DeleteTask)
		}
	}

	engine.GET("/login", userController.LoginIndex)
	engine.GET("/", auth.AuthRequired, userController.Index)

	userRoute := engine.Group("/api/user")
	{
			userRoute.GET("/auth/:provider", userController.Login)
			userRoute.GET("/callback/:provider", userController.Callback)
			userRoute.GET("/logout", userController.Logout)
	}

	engine.GET("/entering", auth.AuthRequired, userController.Entering)
	return engine

}
