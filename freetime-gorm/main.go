package main

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db := db_connect()

	var users users
	users.ID = 8
	// users.Name = "Joko Widodo RI"
	// users.Email = "jokowidodori@gmail.com"
	// users.Password = "234567"
	// users.Address = "IKN"
	// users.Phone = "0866789123456"
	// users.Birth_date = time.Date(2007, 7, 7, 0, 0, 0, 0, time.UTC)

	// db.Create(&users)
	// db.Save(&users)
	// db.Delete(&users)
	read(db)
}

func read(db *gorm.DB) {
	var users []users
	db.Raw("SELECT * FROM users").Scan(&users)
	// db.First(&users)

	for _, value := range users {
		fmt.Println(value)
	}
}

type users struct {
	ID         int
	Name       string
	Email      string
	Password   string
	Address    string
	Phone      string
	Birth_date time.Time
}

func db_connect() *gorm.DB {
	dsn := "host=aws-0-ap-southeast-1.pooler.supabase.com user=postgres.czplcbtzaojfwyvhyzsl password=yytkwE06EO9YDXzR dbname=postgres port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	//user=postgres.czplcbtzaojfwyvhyzsl password=[YOUR-PASSWORD] host=aws-0-ap-southeast-1.pooler.supabase.com port=5432 dbname=postgres

	if err != nil {
		fmt.Println(err)
	}

	return db
}
