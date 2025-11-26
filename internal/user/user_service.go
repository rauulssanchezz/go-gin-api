package user

import (
	"os"

	jwt "github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Create(user User) error
	Update(id string, user User) error
	GetById(id string) (UserResponse, error)
	Delete(id string) error
	GetByEmail(email string) (UserResponse, error)
	Login(email string, password string) (UserResponseLogin, error)
}

type UserServiceStruct struct {
	Repo UserRepositoryStruct
}

func NewUserService(repo UserRepositoryStruct) UserServiceStruct {
	return UserServiceStruct{
		Repo: repo,
	}
}

func (service *UserServiceStruct) Create(user User) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.Password = string(hashed)
	return service.Repo.Create(user)
}

func (service *UserServiceStruct) Update(id string, user User) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.Password = string(hashed)
	return service.Repo.Update(user, id)
}

func (service *UserServiceStruct) Delete(id string) error {
	return service.Repo.Delete(id)
}

func (service *UserServiceStruct) GetByEmail(email string) (UserResponse, error) {
	return service.Repo.GetByEmail(email)
}

func (service *UserServiceStruct) GetById(id string) (UserResponse, error) {
	return service.Repo.GetById(id)
}

func (service *UserServiceStruct) Login(email string, password string) (UserResponseLogin, error) {
	userResponse, err := service.Repo.Login(email, password)

	if err != nil {
		return UserResponseLogin{}, err
	}

	secret := os.Getenv("JWT_SECRET")
	expTime := os.Getenv("EXP_TIME")
	claims := jwt.MapClaims{
		"sub":   userResponse.ID,
		"email": userResponse.Email,
		"exp":   expTime,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))

	if err != nil {
		return UserResponseLogin{}, err
	}

	var userResponseLogin UserResponseLogin = UserResponseLogin{
		userResponse.ID,
		userResponse.Name,
		userResponse.Email,
		signedToken,
	}

	return userResponseLogin, nil
}
