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
	engine.LoadHTMLGlob("./public/index.html")
	engine.Static("/_nuxt", "./public/_nuxt")
	engine.StaticFile("/favicon.ico", "./public/favicon.ico")

	taskRoute := engine.Group("/api/task")
	{
		v1 := taskRoute.Group("/v1")
		{
			v1.POST("/add", auth.AuthRequired, taskController.AddTask)
			v1.GET("/list", auth.AuthRequired, taskController.ListTask)
			v1.PUT("/update/:id", auth.AuthRequired, taskController.UpdateTask)
			v1.DELETE("/delete/:id", auth.AuthRequired, taskController.DeleteTask)
		}
	}
	engine.GET("/", auth.AuthRequired, userController.Index)
	engine.GET("/entering", auth.AuthRequired, userController.Entering)
	engine.GET("/login", userController.LoginIndex)

	userRoute := engine.Group("/api/user")
	{
			userRoute.GET("/auth/:provider", userController.Login)
			userRoute.GET("/callback/:provider", userController.Callback)
			userRoute.GET("/logout", userController.Logout)
		v1 := userRoute.Group("/v1")
		{
			v1.PUT("/update", userController.UpdateUser)
		}
	}

	return engine

}
