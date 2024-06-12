package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Activity string
	Mark     bool
	Owner    uint
}

type TodoModel struct {
	db *gorm.DB
}

func NewTodoModel(connection *gorm.DB) *TodoModel {
	return &TodoModel{
		db: connection,
	}
}

func (add *TodoModel) AddTodo(newData Todo) {
	newData.Mark = false
	err := add.db.Create(&newData).Error
	if err != nil {
		fmt.Println(err)
	}
}

func (read *TodoModel) ReadTodo(user_id uint) []Todo {
	var todo []Todo
	read.db.Where("owner = ?", user_id).Find(&todo)
	return todo
}

func (update *TodoModel) UpdateTodo(user_id uint, todo_id int) {
	update.db.Model(&Todo{}).Where("id = ? AND owner = ?", todo_id, user_id).Update("mark", true)
}

func (delete *TodoModel) DeleteTodo(user_id uint, todo_id int) {
	delete.db.Where("id = ? AND owner = ?", todo_id, user_id).Delete(&Todo{})
}
