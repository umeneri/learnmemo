package service

import (
	"api/domain/model"

	"github.com/go-xorm/xorm"
)

type TaskService interface {
	SetTask(*model.Task) error
	GetTaskList() []model.Task
	UpdateTask(*model.Task) error
	DeleteTask(id int) error
}

type taskService struct {
	dbEngine *xorm.Engine
}

func NewTaskService(dbName string) TaskService {
	dbEngine := initDbEngine(dbName)
	return &taskService{dbEngine}
}

func (t *taskService) SetTask(task *model.Task) error {
	_, err := t.dbEngine.Insert(task)
	if err != nil {
		return err
	}
	return nil
}

func (t *taskService) GetTaskList() []model.Task {
	tasks := make([]model.Task, 0)
	err := t.dbEngine.Limit(10, 0).Find(&tasks)
	if err != nil {
		panic(err)
	}
	return tasks
}

func (t *taskService) UpdateTask(newTask *model.Task) error {
	_, err := t.dbEngine.Id(newTask.Id).Update(newTask)
	if err != nil {
		return err
	}
	return nil
}

func (t *taskService) DeleteTask(id int) error {
	task := new(model.Task)
	_, err := t.dbEngine.Id(id).Delete(task)
	if err != nil {
		return err
	}
	return nil
}
