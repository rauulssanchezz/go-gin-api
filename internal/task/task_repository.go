package task

import (
	"database/sql"
)

type TaskRepository interface {
	Create(task Task) error
	Update(id string, task Task) error
	GetAll() ([]Task, error)
	Delete(id string) error
	GetById(id string) (Task, error)
}

type TaskRepositoryStruct struct {
	DB *sql.DB
}

func NewTaskRepositoryStruct(db *sql.DB) TaskRepositoryStruct {
	return TaskRepositoryStruct{DB: db}
}

func (repository *TaskRepositoryStruct) Create(task Task) error {
	const query string = `INSERT INTO tasks(title, description, done)
							VALUES ($1, $2, $3)`

	_, err := repository.DB.Exec(query, task.Title, task.Description, task.Done)

	return err
}

func (repository *TaskRepositoryStruct) Update(id string, task Task) error {
	const query string = `UPDATE tasks SET title = $1, description = $2, done = $3 where id = $4`

	res, err := repository.DB.Exec(query, task.Title, task.Description, task.Done, id)

	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected <= 0 {
		return sql.ErrNoRows
	}

	return err
}

func (repository *TaskRepositoryStruct) GetAll() ([]Task, error) {
	const query string = `SELECT * FROM tasks`

	rows, err := repository.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Done, &task.CreatedAt)

		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (repository *TaskRepositoryStruct) Delete(id string) error {
	const query string = `DELETE FROM tasks WHERE id = $1`

	res, err := repository.DB.Exec(query, id)

	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected <= 0 {
		return sql.ErrNoRows
	}

	return err
}

func (repository *TaskRepositoryStruct) GetById(id string) (Task, error) {
	const query string = `SELECT * FROM tasks WHERE id = $1`

	var task Task
	err := repository.DB.QueryRow(query, id).Scan(&task.ID, &task.Title, &task.Description, &task.Done, &task.CreatedAt)

	if err != nil {
		return Task{}, err
	}

	return task, nil
}
