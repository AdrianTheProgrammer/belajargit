package main

import (
	"fmt"
	"github/configs"
	"github/internal/controllers"
	"github/internal/models"
)

func main() {
	setup := configs.ImportSetting()
	connection, err := configs.ConnectDB(setup)

	if err != nil {
		fmt.Println("Stop program, masalah pada database", err.Error())
		return
	}

	// connection.AutoMigrate(&models.User{}, &models.Todo{})

	um := models.NewUserModel(connection)
	uc := controllers.NewUserController(um)

	tu := models.NewTodoModel(connection)
	tc := controllers.NewTodoController(tu)

	main_menu(uc, tc)
}

func main_menu(uc *controllers.UserController, tc *controllers.TodoController) {
	var inputMenu int
	for inputMenu != 9 {
		fmt.Println("Pilih Menu")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("9. Keluar")
		fmt.Print("Masukkan Input: ")
		fmt.Scanln(&inputMenu)
		if inputMenu == 1 {
			fmt.Println()
			login(uc, tc)
		} else if inputMenu == 2 {
			fmt.Println()
			uc.Register()
		}
	}

	fmt.Println("Program Berakhir. Terima Kasih!")
}

func login(uc *controllers.UserController, tc *controllers.TodoController) {
	var isLogin = true
	data_login, err := uc.Login()
	if err != nil {
		fmt.Println("Terjadi error pada saat login, error: ", err.Error())
		return
	}

	for isLogin {
		isLogin = todo_menu(tc, data_login)
	}
}

func todo_menu(tc *controllers.TodoController, data_login models.User) bool {
	var input_todo int

	fmt.Printf("Selamat datang %s!\n", data_login.Name)
	tc.ReadTodo(data_login.ID)

	fmt.Println("Pilih menu")
	fmt.Println("1. Tambah Kegiatan | 2. Perbarui Kegiatan | 3. Hapus Kegiatan | 9. Logout")
	fmt.Print("Masukkan Input: ")
	fmt.Scanln(&input_todo)
	fmt.Println()

	var isLogin bool

	if input_todo == 9 {
		isLogin = false
	} else if input_todo == 1 {
		tc.AddTodo(data_login.ID)
		isLogin = true
	} else if input_todo == 2 {
		tc.UpdateTodo(data_login.ID)
		isLogin = true
	} else if input_todo == 3 {
		tc.DeleteTodo(data_login.ID)
		isLogin = true
	}

	return isLogin
}
