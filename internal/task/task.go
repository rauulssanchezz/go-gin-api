package task

import "time"

type Task struct {
	ID          int64     `json:"id" db:"id"`
	Title       string    `json:"title" db:"title" validate:"required"`
	Description string    `json:"description" db:"description" validate:"required"`
	Done        bool      `json:"done" db:"done"`
	CreatedAt   time.Time `json:"created_at db:"created_at`
}
