package user

import (
	"database/sql"
)

type UserRepository interface {
	Create(user User) error
	Update(user User) error
	GetById(id string) (User, error)
	Delete(id string) error
	GetByEmail(email string) (User, error)
}

type UserRepositoryStruct struct {
	DB *sql.DB
}

func NewUserRepositoryStruct(DB *sql.DB) *UserRepositoryStruct {
	return &UserRepositoryStruct{DB: DB}
}

func (repository *UserRepositoryStruct) Create(user User) error {
	const query string = `INSERT INTO users(name, email, password) VALUES($1, $2, $3)`

	_, err := repository.DB.Exec(query, user.Name, user.Email, user.Password)

	return err
}

func (repository *UserRepositoryStruct) Update(user User, id string) error {
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

func (repository *UserRepositoryStruct) GetById(id string) (User, error) {
	const query string = `SELECT id, name, email, password FROM users WHERE id = $1`

	var user User
	err := repository.DB.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	if err != nil {
		return User{}, err
	}

	return user, err
}

func (repository *UserRepositoryStruct) GetByEmail(email string) (User, error) {
	const query string = `SELECT * FROM users WHERE email = WHERE id = $1`

	var user User
	err := repository.DB.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	if err != nil {
		return User{}, err
	}

	return user, err
}
