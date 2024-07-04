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

func ToAllTodosEntity(alltodos []Todos) []todos.Todos {
	var result []todos.Todos

	for _, val := range alltodos {
		result = append(result, todos.Todos{
			ID:       val.ID,
			Activity: val.Activity,
			Date:     val.Date,
			Status:   val.Status,
			UserID:   val.UserID,
		})
	}

	return result
}

func ToTodosData(input todos.Todos) Todos {
	return Todos{
		Activity: input.Activity,
		Date:     input.Date,
		Status:   input.Status,
		UserID:   input.UserID,
	}
}
