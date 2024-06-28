package todos

import (
	"github/internal/models"
	"time"
)

type AllTodos struct {
	ID       uint      `json:"id"`
	Activity string    `json:"activity"`
	Date     time.Time `json:"date"`
	Status   bool      `json:"status"`
	UserID   uint      `json:"user_id"`
}

func ToAllTodos(todos []models.Todos) []AllTodos {
	var result []AllTodos

	for _, val := range todos {
		result = append(result, AllTodos{
			ID:       val.ID,
			Activity: val.Activity,
			Date:     val.Date,
			Status:   val.Status,
			UserID:   val.UserID,
		})
	}

	return result
}
