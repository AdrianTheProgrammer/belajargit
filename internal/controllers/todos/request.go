package todos

import (
	"github/internal/models"
	"time"
)

type TodosRequest struct {
	Activity string    `json:"activity"`
	Date     time.Time `json:"date"`
}

func ToModelTodos(tr TodosRequest, UserID uint) models.Todos {
	return models.Todos{
		Activity: tr.Activity,
		Date:     tr.Date,
		UserID:   UserID,
	}
}
