package repository

import (
	"api/domain/model"
	"api/domain/repository"

	"github.com/go-xorm/xorm"
)

type taskRepository struct {
	dbEngine *xorm.Engine
}

func NewTaskRepository(dbName string) repository.TaskRepository {
	dbEngine := initDbEngine(dbName)
	return &taskRepository{dbEngine}
}

func (t *taskRepository) SetTask(task *model.Task) error {
	_, err := t.dbEngine.Insert(task)
	if err != nil {
		return err
	}
	return nil
}

func (t *taskRepository) GetTaskList() []model.Task {
	tasks := make([]model.Task, 0)
	err := t.dbEngine.Limit(20, 0).Find(&tasks)
	if err != nil {
		panic(err)
	}
	return tasks
}

func (t *taskRepository) UpdateTask(newTask *model.Task) error {
	_, err := t.dbEngine.Id(newTask.Id).Update(newTask)
	if err != nil {
		return err
	}
	return nil
}

func (t *taskRepository) DeleteTask(id int) error {
	task := new(model.Task)
	_, err := t.dbEngine.Id(id).Delete(task)
	if err != nil {
		return err
	}
	return nil
}
