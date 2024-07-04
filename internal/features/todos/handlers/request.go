package handlers

import (
	"github/internal/features/todos/repositories"
	"time"
)

type TodosRequest struct {
	Activity string    `json:"activity"`
	Date     time.Time `json:"date"`
	Status   bool
}

func ToRepoTodos(tr TodosRequest, UserID uint) repositories.Todos {
	return repositories.Todos{
		Activity: tr.Activity,
		Date:     tr.Date,
		Status:   tr.Status,
		UserID:   UserID,
	}
}
