package repository

import (
	"api/domain/model"
	"api/domain/repository"

	"github.com/go-xorm/xorm"
	"github.com/lunny/log"
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

func (t *taskRepository) GetTaskList(userId int64) []model.Task {
	tasks := make([]model.Task, 0)
	err := t.dbEngine.Where("user_id = ?", userId).Limit(20, 0).Find(&tasks)
	if err != nil {
		panic(err)
	}
	log.Println("tasks list")
	log.Println(tasks)
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
