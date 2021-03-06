package router

import (
	"api/interfaces/auth"
	"api/interfaces/controller"
	"os"

	"api/interfaces/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoute(taskController controller.TaskController, userController controller.UserController, healthController controller.HealthController) *gin.Engine {
	engine := gin.Default()

	if os.Getenv("ENV") == "dev" || os.Getenv("ENV") == "test" {
		engine.Use(middleware.PrintAccessLog)
	}
	engine.LoadHTMLGlob("./public/index.html")
	engine.Static("/_nuxt", "./public/_nuxt")
	engine.StaticFile("/favicon.ico", "./public/favicon.ico")

	engine.GET("/health", healthController.Index)

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
	engine.GET("/", auth.AuthRequiredPage, userController.Index)
	engine.GET("/entering", auth.AuthRequiredPage, userController.Entering)
	engine.GET("/login", userController.LoginIndex)

	userRoute := engine.Group("/api/user")
	{
		userRoute.GET("/simple-login", userController.SimpleLogin)
		userRoute.GET("/auth/:provider", userController.Login)
		userRoute.GET("/callback/:provider", userController.Callback)
		userRoute.GET("/logout", userController.Logout)
		v1 := userRoute.Group("/v1")
		{
			v1.PUT("/update", auth.AuthRequired, userController.UpdateUser)
		}
	}

	return engine

}
