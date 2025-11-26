package user

type User struct {
	ID       string `json:"id" db:"id"`
	Name     string `json:"title" db:"title"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserResponseLogin struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}
