package user

import (
	"database/sql"
)

type UserRepository interface {
	Create(user User) error
	Update(user User, id string) error
	GetById(id string) (UserResponse, error)
	Delete(id string) error
	GetByEmail(email string) (UserResponse, error)
	Login(email string, password string) (UserResponse, error)
}

type UserRepositoryStruct struct {
	DB *sql.DB
}

func NewUserRepository(DB *sql.DB) UserRepositoryStruct {
	return UserRepositoryStruct{DB: DB}
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

func (repository *UserRepositoryStruct) GetById(id string) (UserResponse, error) {
	const query string = `SELECT id, name, email, password FROM users WHERE id = $1`

	var user UserResponse
	err := repository.DB.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email)

	if err != nil {
		return UserResponse{}, err
	}

	return user, err
}

func (repository *UserRepositoryStruct) GetByEmail(email string) (UserResponse, error) {
	const query string = `SELECT * FROM users WHERE email = $1`

	var user UserResponse
	err := repository.DB.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email)

	if err != nil {
		return UserResponse{}, err
	}

	return user, err
}

func (repository *UserRepositoryStruct) Login(email string, password string) (UserResponse, error) {
	const query string = `SELECT * FROM users WHERE email = $1 and password $2`

	var user UserResponse
	err := repository.DB.QueryRow(query, email, password).Scan(&user.ID, &user.Name, &user.Email)

	if err != nil {
		return UserResponse{}, err
	}

	return user, err
}

func (repository *UserRepositoryStruct) Delete(id string) error {
	const query string = `DELETE FROM users WHERE id = $1`

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
