package usecase

import (
	"api/domain/model"
	"api/domain/repository"
)

type TaskUseCase interface {
	SetTask(*model.Task) error
	GetTaskList() []model.Task
	UpdateTask(*model.Task) error
	DeleteTask(id int) error
	GetUserId() int64
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

func (t *taskUseCase) GetTaskList() []model.Task {
	return t.taskRepository.GetTaskList()
}

func (t *taskUseCase) UpdateTask(task *model.Task) error {
	return t.taskRepository.UpdateTask(task)
}

func (t *taskUseCase) DeleteTask(id int) error {
	return t.taskRepository.DeleteTask(id)
}

func (t *taskUseCase) GetUserId() int64 {
	return 1
}