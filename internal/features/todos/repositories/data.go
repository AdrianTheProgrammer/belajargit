package repositories

import (
	"github/internal/features/todos"
	"time"

	"gorm.io/gorm"
)

type Todos struct {
	gorm.Model
	Activity string    `json:"activity"`
	Date     time.Time `json:"date"`
	Status   bool      `json:"status"`
	UserID   uint      `json:"user_id"`
}

func ToTodosEntity(input Todos) todos.Todos {
	return todos.Todos{
		ID:       input.ID,
		Activity: input.Activity,
		Date:     input.Date,
		Status:   input.Status,
		UserID:   input.UserID,
	}
}

func ToTodosData(input todos.Todos) Todos {
	return Todos{
		Activity: input.Activity,
		Date:     input.Date,
		Status:   input.Status,
		UserID:   input.UserID,
	}
}
