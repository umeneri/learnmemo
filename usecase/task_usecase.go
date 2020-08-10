package usecase

import (
	"api/domain/model"
	"api/domain/repository"
)

type TaskUseCase interface {
	SetTask(*model.Task) error
	GetTaskList(int64) []model.Task
	UpdateTask(*model.Task) error
	DeleteTask(id int) error
}

type taskUseCase struct {
	taskRepository repository.TaskRepository
}

func NewTaskUseCase(taskRepository repository.TaskRepository) TaskUseCase {
  return &taskUseCase{
		taskRepository: taskRepository,
	}
}

func (t *taskUseCase) SetTask(task *model.Task) error {
	return t.taskRepository.SetTask(task)
}

func (t *taskUseCase) GetTaskList(userId int64) []model.Task {
	return t.taskRepository.GetTaskList(userId)
}

func (t *taskUseCase) UpdateTask(task *model.Task) error {
	return t.taskRepository.UpdateTask(task)
}

func (t *taskUseCase) DeleteTask(id int) error {
	return t.taskRepository.DeleteTask(id)
}