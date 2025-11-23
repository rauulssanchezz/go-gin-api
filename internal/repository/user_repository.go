package repository

import (
	"database/sql"

	"github.com/rauulssanchezz/go-gin-api/internal/model"
)

type UserRepository interface {
	Create(user *model.User) error
	Update(user *model.User) error
	GetById(id string) (model.User, error)
	Delete(id string) error
}

type UserRepositoryStruct struct {
	DB *sql.DB
}

func NewUserRepositoryStruct(DB *sql.DB) *UserRepositoryStruct {
	return &UserRepositoryStruct{DB: DB}
}

func (repository *UserRepositoryStruct) Create(user model.User) error {
	const query string = `INSERT INTO users(name, email, password) VALUES($1, $2, $3)`

	_, err := repository.DB.Exec(query, user.Name, user.Email, user.Password)

	return err
}

func (repository *UserRepositoryStruct) Update(user model.User, id string) error {
	const query string = `UPDATE users SET name = $1, email = $2, password = $3 where id = $4`

	res, err := repository.DB.Exec(query, user.Name, user.Email, user.Password, id)

	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()

	if rowsAffected <= 0 {
		return sql.ErrNoRows
	}

	return err
}

func (repository *UserRepositoryStruct) GetById(id string) (model.User, error) {
	const query string = `SELECT id, name, email, password FROM users`

	var user model.User
	err := repository.DB.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, user.Password)

	if err != nil {
		return model.User{}, err
	}

	return user, err
}
