package controller

import (
	"api/domain/model"
	"api/interfaces/auth"
	"api/usecase"
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
	Title       string `json:"title" binding:"required"`
	ElapsedTime int64  `json:"elapsedTime" binding:"min=0"`
	Status      int    `json:"status" binding:"min=0"`
}

func NewTaskController(useCase usecase.TaskUseCase) TaskController {
	return &taskController{
		taskUseCase: useCase,
	}
}

func (t *taskController) AddTask(c *gin.Context) {
	taskForm := TaskForm{}
	err := c.BindJSON(&taskForm)

	userId, err := auth.GetUserId(c)
	log.Println(userId)
	log.Println(err)

	task := model.Task{
		UserId:      userId,
		Title:       taskForm.Title,
		ElapsedTime: taskForm.ElapsedTime,
		Status:      taskForm.Status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
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
	userId, err := auth.GetUserId(c)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	taskLists := t.taskUseCase.GetTaskList(userId)
	c.JSON(http.StatusOK, gin.H{
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
		Id:          intId,
		Title:       taskForm.Title,
		ElapsedTime: taskForm.ElapsedTime,
		Status:      taskForm.Status,
		UpdatedAt:   time.Now(),
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
