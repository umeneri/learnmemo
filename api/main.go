package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"api/controller"
	"api/middleware"
)

func main() {
	engine := gin.Default()
	engine.Use(middleware.RecordUaAndTime)
	taskEngine := engine.Group("/task")
	{
		v1 := taskEngine.Group("/v1")
		{
			v1.POST("/add", controller.TaskAdd)
			v1.GET("/list", controller.TaskList)
			v1.PUT("/update/:id", controller.TaskUpdate)
			v1.DELETE("/delete/:id", controller.TaskDelete)
		}
	}
	engine.Run(":3000")
}
