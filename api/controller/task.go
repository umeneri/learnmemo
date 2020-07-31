package controller

import (
	"api/model"
	"api/service"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type TaskForm struct {
	Title          string `json:"title" binding:"required"`
	ProgressMinute int64  `json:"progressMinute" binding:"required"`
	Status         int    `json:"status" binding:"required"`
}

func GetUserId(c *gin.Context) int64 {
	return 1
}

func TaskAdd(c *gin.Context) {
	taskForm := TaskForm{}
	err := c.Bind(&taskForm)
	log.Println(taskForm)
	log.Println(err)

	task := model.Task{
		UserId:         GetUserId(c),
		Title:          taskForm.Title,
		ProgressMinute: taskForm.ProgressMinute,
		Status:         taskForm.Status,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	log.Println(c.ContentType())

	taskService := service.TaskService{}
	err = taskService.SetTask(&task)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

func TaskList(c *gin.Context) {
	taskService := service.TaskService{}
	taskLists := taskService.GetTaskList()
	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"data":    taskLists,
	})
}

func TaskUpdate(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	taskForm := TaskForm{}
	err = c.Bind(&taskForm)
	log.Println(taskForm)
	log.Println(err)

	task := model.Task{
		Id:             intId,
		Title:          taskForm.Title,
		ProgressMinute: taskForm.ProgressMinute,
		Status:         taskForm.Status,
		UpdatedAt: time.Now(),
	}

	if err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	taskService := service.TaskService{}
	err = taskService.UpdateTask(&task)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

func TaskDelete(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	taskService := service.TaskService{}
	err = taskService.DeleteTask(int(intId))
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}
