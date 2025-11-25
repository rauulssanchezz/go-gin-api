package user

type User struct {
	ID       string `json:"id" db:"id"`
	Name     string `json:"title" db:"title"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}
