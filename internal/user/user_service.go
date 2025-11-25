package user

type UserService interface {
	Create(user User) error
	Update(id string, user User) error
	GetById(id string) (User, error)
	Delete(id string) error
	GetByEmail(email string) (User, error)
	Login(email string, password string) (User, error)
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
	return service.Repo.Create(user)
}

func (service *UserServiceStruct) Update(id string, user User) error {
	return service.Repo.Update(user, id)
}

func (service *UserServiceStruct) Delete(id string) error {
	return service.Repo.Delete(id)
}

func (service *UserServiceStruct) GetByEmail(email string) (User, error) {
	return service.Repo.GetByEmail(email)
}

func (service *UserServiceStruct) GetById(id string) (User, error) {
	return service.Repo.GetById(id)
}

func (service *UserServiceStruct) Login(email string, password string) (User, error) {
	return service.Repo.Login(email, password)
}
