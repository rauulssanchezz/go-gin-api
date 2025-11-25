package user

type UserService interface {
	Create(user User) error
	Update(id string, user User) error
	GetById(id string) (User, error)
	Delete(id string) error
	GetByEmail(email string) (User, error)
}

type UserServiceStruct struct {
	Repo UserRepository
}

func NewUserService(repo UserRepository) *UserServiceStruct {
	return &UserServiceStruct{
		Repo: repo,
	}
}

func (service *UserServiceStruct) Create(user User) error {
	service.Repo.Create(user)
}
