package services

import (
	"github/internal/features/todos"
)

type TodosServices struct {
	qry todos.Query
}

func NewTodosSer(q todos.Query) todos.Services {
	return &TodosServices{
		qry: q,
	}
}

func (ts *TodosServices) CreateTodo(todo todos.Todos) error {
	return ts.qry.CreateTodo(todo)
}

func (ts *TodosServices) ReadAllTodos(userID uint) ([]todos.Todos, error) {
	return ts.qry.ReadAllTodos(userID)
}

func (ts *TodosServices) UpdateTodo(todoID uint, todo todos.Todos) error {
	todo.ID = todoID
	return ts.qry.UpdateTodo(todo)
}

func (ts *TodosServices) DeleteTodo(todoID uint) error {
	return ts.qry.DeleteTodo(todoID)
}
