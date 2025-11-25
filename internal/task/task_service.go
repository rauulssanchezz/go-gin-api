package task

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type TaskService interface {
	Create(Task) error
	Update(Task) error
	GetAll() ([]Task, error)
	GetById(id string) (Task, error)
	Delete(id string) error
}

type TaskServiceStruct struct {
	Repo TaskRepository
}

func NewTaskService(repo TaskRepository) *TaskServiceStruct {
	return &TaskServiceStruct{
		Repo: repo,
	}
}

func (TaskServiceStruct *TaskServiceStruct) Create(task Task) error {
	err := validate.Struct(task)

	if err != nil {
		return err
	}

	return TaskServiceStruct.Repo.Create(task)
}

func (TaskServiceStruct *TaskServiceStruct) Update(id string, task Task) error {
	err := validate.Struct(task)

	if err != nil {
		return err
	}

	return TaskServiceStruct.Repo.Update(id, task)
}

func (TaskServiceStruct *TaskServiceStruct) GetAll() ([]Task, error) {
	return TaskServiceStruct.Repo.GetAll()
}

func (TaskServiceStruct *TaskServiceStruct) GetById(id string) (Task, error) {
	return TaskServiceStruct.Repo.GetById(id)
}

func (TaskServiceStruct *TaskServiceStruct) Delete(id string) error {
	return TaskServiceStruct.Repo.Delete(id)
}
