package repository

import (
	"database/sql"

	"github.com/rauulssanchezz/go-gin-api/internal/model"
)

type TaskRepository interface {
	Create(task *model.Task) error
	Update(id string, task *model.Task) error
	GetAll() ([]model.Task, error)
	Delete(id string) error
	GetById(id string) (model.Task, error)
}

type PostgreSQLTaskRepository struct {
	DB *sql.DB
}

func NewPostgreSQLTaskRepository(db *sql.DB) *PostgreSQLTaskRepository {
	return &PostgreSQLTaskRepository{DB: db}
}

func (repository *PostgreSQLTaskRepository) Create(task *model.Task) error {
	const query string = `INSERT INTO tasks(title, description, done)
							VALUES ($1, $2, $3)`

	_, err := repository.DB.Exec(query, task.Title, task.Description, task.Done)

	return err
}

func (repository *PostgreSQLTaskRepository) Update(id string, task *model.Task) error {
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

func (repository *PostgreSQLTaskRepository) GetAll() ([]model.Task, error) {
	const query string = `SELECT * FROM tasks`

	rows, err := repository.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var task model.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Done, &task.CreatedAt)

		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (repository *PostgreSQLTaskRepository) Delete(id string) error {
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

func (repository *PostgreSQLTaskRepository) GetById(id string) (model.Task, error) {
	const query string = `SELECT * FROM tasks WHERE id = $1`

	var task model.Task
	err := repository.DB.QueryRow(query, id).Scan(&task.ID, &task.Title, &task.Description, &task.Done, &task.CreatedAt)

	if err != nil {
		return model.Task{}, err
	}

	return task, nil
}
