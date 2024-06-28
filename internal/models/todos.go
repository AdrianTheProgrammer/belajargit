package models

import (
	"time"

	"gorm.io/gorm"
)

type TodosMod struct {
	db *gorm.DB
}

func NewTodosMod(connection *gorm.DB) *TodosMod {
	return &TodosMod{
		db: connection,
	}
}

type Todos struct {
	gorm.Model
	Activity string    `json:"activity"`
	Date     time.Time `json:"date"`
	Status   bool      `json:"status"`
	UserID   uint      `json:"user_id"`
}

func (um *TodosMod) CreateTodo(todo Todos) error {
	err := um.db.Create(&todo).Error

	if err == nil {
		return nil
	}

	return err
}

func (um *TodosMod) ReadAllTodos(userID string) ([]Todos, error) {
	var todos []Todos
	err := um.db.Where("user_id = ?", userID).Find(&todos).Error

	if err != nil {
		return []Todos{}, err
	}

	return todos, nil
}

func (um *TodosMod) UpdateTodo(todo Todos) error {
	err := um.db.Save(&todo).Error

	if err != nil {
		return err
	}

	return nil
}

func (um *TodosMod) DeleteTodo(id string) error {
	err := um.db.Delete(&Todos{}, id).Error

	if err != nil {
		return err
	}

	return nil
}
