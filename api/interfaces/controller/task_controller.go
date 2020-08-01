package controller

import (
	"api/domain/model"
	"api/usecase"
	"api/interfaces/auth"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type TaskController interface {
	AddTask(c *gin.Context)
	ListTask(c *gin.Context)
	UpdateTask(c *gin.Context)
	DeleteTask(c *gin.Context)
}

type taskController struct {
	taskUseCase usecase.TaskUseCase
}

type TaskForm struct {
	Title          string `json:"title" binding:"required"`
	ProgressMinute int64  `json:"progressMinute" binding:"required"`
	Status         int    `json:"status" binding:"required"`
}

func NewTaskController(useCase usecase.TaskUseCase) TaskController {
	return &taskController{
		taskUseCase: useCase,
	}
}

func (t *taskController) AddTask(c *gin.Context) {
	taskForm := TaskForm{}
	err := c.BindJSON(&taskForm)
	log.Println(taskForm)
	log.Println(err)

	task := model.Task{
		UserId:         auth.GetUserId(),
		Title:          taskForm.Title,
		ProgressMinute: taskForm.ProgressMinute,
		Status:         taskForm.Status,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	err = t.taskUseCase.SetTask(&task)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

func (t *taskController) ListTask(c *gin.Context) {
	taskLists := t.taskUseCase.GetTaskList()
	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"data":    taskLists,
	})
}

func (t *taskController) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	taskForm := TaskForm{}
	err = c.BindJSON(&taskForm)
	log.Println(taskForm)
	log.Println(err)

	task := model.Task{
		Id:             intId,
		Title:          taskForm.Title,
		ProgressMinute: taskForm.ProgressMinute,
		Status:         taskForm.Status,
		UpdatedAt:      time.Now(),
	}

	if err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	err = t.taskUseCase.UpdateTask(&task)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

func (t *taskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	err = t.taskUseCase.DeleteTask(int(intId))
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}
