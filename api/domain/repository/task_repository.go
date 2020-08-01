package repository

import "api/domain/model"

type TaskRepository interface {
	SetTask(*model.Task) error
	GetTaskList() []model.Task
	UpdateTask(*model.Task) error
	DeleteTask(id int) error
}

