package controllers

import (
	"fmt"
	"github/internal/models"
)

type TodoController struct {
	model *models.TodoModel
}

func NewTodoController(m *models.TodoModel) *TodoController {
	return &TodoController{
		model: m,
	}
}

func (tc *TodoController) AddTodo(user_id uint) {
	var newData models.Todo
	fmt.Print("Masukkan Aktivitas: ")
	fmt.Scanln(&newData.Activity)
	newData.Owner = user_id
	tc.model.AddTodo(newData)
}

func (tc *TodoController) ReadTodo(user_id uint) {
	var todo_list []models.Todo = tc.model.ReadTodo(user_id)
	if len(todo_list) != 0 {
		fmt.Println("=== DAFTAR KEGIATAN ANDA ===")

		for _, val := range todo_list {
			var mark_as string
			if !val.Mark {
				mark_as = "Belum Selesai"
			} else {
				mark_as = "Selesai"
			}
			fmt.Printf("%d | %s | %s\n", val.ID, val.Activity, mark_as)
		}
	} else {
		fmt.Println("=== ANDA BELUM MEMILIKI KEGIATAN ===")
	}
	fmt.Println()
}

func (tc *TodoController) UpdateTodo(user_id uint) {
	var todo_id int
	fmt.Println("=== PERBARUI KEGIATAN ===")
	fmt.Println("Masukkan '0' untuk membatalkan pembaruan.")
	fmt.Print("Masukkan ID Kegiatan yang ingin anda tandai sebagai 'Selesai': ")
	fmt.Scanln(&todo_id)
	fmt.Println()

	if todo_id != 0 {
		tc.model.UpdateTodo(user_id, todo_id)
	}
}

func (tc *TodoController) DeleteTodo(user_id uint) {
	var todo_id int
	fmt.Println("=== HAPUS KEGIATAN ===")
	fmt.Println("Masukkan '0' untuk membatalkan penghapusan.")
	fmt.Print("Masukkan ID Kegiatan yang akan dihapus: ")
	fmt.Scanln(&todo_id)
	fmt.Println()

	if todo_id != 0 {
		tc.model.DeleteTodo(user_id, todo_id)
	}
}
