package repositories

import (
	"github/internal/features/todos"

	"gorm.io/gorm"
)

type TodosQue struct {
	db *gorm.DB
}

func NewTodosQue(connection *gorm.DB) *TodosQue {
	return &TodosQue{
		db: connection,
	}
}

func (tq *TodosQue) CreateTodo(todo todos.Todos) error {
	todocnv := ToTodosData(todo)
	err := tq.db.Create(&todocnv).Error

	if err != nil {
		return err
	}

	return nil
}

func (tq *TodosQue) ReadAllTodos(userID uint) ([]todos.Todos, error) {
	var alltodos []Todos
	err := tq.db.Where("user_id = ?", userID).Find(&alltodos).Error

	if err != nil {
		return []todos.Todos{}, err
	}

	return ToAllTodosEntity(alltodos), nil
}

func (tq *TodosQue) UpdateTodo(todo todos.Todos) error {
	todocnv := ToTodosData(todo)
	err := tq.db.Save(&todocnv).Error

	if err != nil {
		return err
	}

	return nil
}

func (tq *TodosQue) DeleteTodo(id uint) error {
	err := tq.db.Delete(&Todos{}, id).Error

	if err != nil {
		return err
	}

	return nil
}
