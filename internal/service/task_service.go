package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/rauulssanchezz/go-gin-api/internal/model"
	"github.com/rauulssanchezz/go-gin-api/internal/repository"
)

var validate = validator.New()

type TaskService interface {
	Create(*model.Task) error
	Update(*model.Task) error
	GetAll() ([]model.Task, error)
	GetById(id string) (model.Task, error)
	Delete(id string) error
}

type TaskServiceStruct struct {
	Repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskServiceStruct {
	return &TaskServiceStruct{
		Repo: repo,
	}
}

func (TaskServiceStruct *TaskServiceStruct) Create(task *model.Task) error {
	err := validate.Struct(task)

	if err != nil {
		return err
	}

	return TaskServiceStruct.Repo.Create(task)
}

func (TaskServiceStruct *TaskServiceStruct) Update(id string, task *model.Task) error {
	err := validate.Struct(task)

	if err != nil {
		return err
	}

	return TaskServiceStruct.Repo.Update(id, task)
}

func (TaskServiceStruct *TaskServiceStruct) GetAll() ([]model.Task, error) {
	return TaskServiceStruct.Repo.GetAll()
}

func (TaskServiceStruct *TaskServiceStruct) GetById(id string) (model.Task, error) {
	return TaskServiceStruct.Repo.GetById(id)
}

func (TaskServiceStruct *TaskServiceStruct) Delete(id string) error {
	return TaskServiceStruct.Repo.Delete(id)
}
