package services_test

import (
	"github/internal/features/todos"
	"github/internal/features/todos/services"
	"github/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreateTodo(t *testing.T) {
	qry := mocks.NewTQuery(t)
	srv := services.NewTodosSer(qry)

	input := todos.Todos{
		ID:       1,
		Activity: "Lari Pagi",
		Date:     time.Date(2024, 07, 06, 07, 30, 0, 0, time.UTC),
		Status:   false,
		UserID:   1,
	}

	t.Run("Query Error", func(t *testing.T) {
		qry.On("CreateTodo", input).Return(gorm.ErrInvalidData).Once()
		err := srv.CreateTodo(input)

		qry.AssertExpectations(t)
		assert.Error(t, err)
	})

	t.Run("Query Success", func(t *testing.T) {
		qry.On("CreateTodo", input).Return(nil).Once()
		err := srv.CreateTodo(input)

		qry.AssertExpectations(t)
		assert.Nil(t, err)
	})
}

func TestReadAllTodos(t *testing.T) {
	qry := mocks.NewTQuery(t)
	srv := services.NewTodosSer(qry)

	var result []todos.Todos
	result = append(result, todos.Todos{
		ID:       1,
		Activity: "Lari Pagi",
		Date:     time.Date(2024, 07, 06, 07, 30, 0, 0, time.UTC),
		Status:   false,
		UserID:   1,
	})
	result = append(result, todos.Todos{
		ID:       2,
		Activity: "Lari Siang",
		Date:     time.Date(2024, 07, 06, 13, 0, 0, 0, time.UTC),
		Status:   false,
		UserID:   1,
	})

	var userID uint = 1

	t.Run("Query Error", func(t *testing.T) {
		qry.On("ReadAllTodos", userID).Return([]todos.Todos{}, gorm.ErrInvalidData).Once()
		_, err := srv.ReadAllTodos(userID)

		qry.AssertExpectations(t)
		assert.Error(t, err)
	})

	t.Run("Query Success", func(t *testing.T) {
		qry.On("ReadAllTodos", userID).Return(result, nil).Once()
		_, err := srv.ReadAllTodos(userID)

		qry.AssertExpectations(t)
		assert.Nil(t, err)
	})
}

func TestUpdateTodo(t *testing.T) {
	qry := mocks.NewTQuery(t)
	srv := services.NewTodosSer(qry)

	var parameter uint = 1
	input := todos.Todos{
		ID:       parameter,
		Activity: "Lari Malam",
		Date:     time.Date(2024, 07, 06, 20, 0, 0, 0, time.UTC),
		Status:   false,
		UserID:   1,
	}

	t.Run("Query Error", func(t *testing.T) {
		qry.On("UpdateTodo", parameter, input).Return(gorm.ErrInvalidData).Once()
		err := srv.UpdateTodo(parameter, input)

		qry.AssertExpectations(t)
		assert.Error(t, err)
	})

	t.Run("Query Success", func(t *testing.T) {
		qry.On("UpdateTodo", parameter, input).Return(nil).Once()
		err := srv.UpdateTodo(parameter, input)

		qry.AssertExpectations(t)
		assert.Nil(t, err)
	})
}

func TestDeleteTodo(t *testing.T) {
	qry := mocks.NewTQuery(t)
	srv := services.NewTodosSer(qry)

	var parameter uint = 1

	t.Run("Query Error", func(t *testing.T) {
		qry.On("DeleteTodo", parameter).Return(gorm.ErrInvalidData).Once()
		err := srv.DeleteTodo(parameter)

		qry.AssertExpectations(t)
		assert.Error(t, err)
	})

	t.Run("Query Success", func(t *testing.T) {
		qry.On("DeleteTodo", parameter).Return(nil).Once()
		err := srv.DeleteTodo(parameter)

		qry.AssertExpectations(t)
		assert.Nil(t, err)
	})
}
