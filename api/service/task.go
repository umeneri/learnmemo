package service

import (
	"api/model"
)

type TaskService struct{}

func (TaskService) SetTask(task *model.Task) error {
	_, err := DbEngine.Insert(task)
	if err != nil {
		return err
	}
	return nil
}

func (TaskService) GetTaskList() []model.Task {
	tasks := make([]model.Task, 0)
	err := DbEngine.Limit(10, 0).Find(&tasks)
	if err != nil {
		panic(err)
	}
	return tasks
}

func (TaskService) UpdateTask(newTask *model.Task) error {
	_, err := DbEngine.Id(newTask.Id).Update(newTask)
	if err != nil {
		return err
	}
	return nil
}

func (TaskService) DeleteTask(id int) error {
	task := new(model.Task)
	_, err := DbEngine.Id(id).Delete(task)
	if err != nil {
		return err
	}
	return nil
}
