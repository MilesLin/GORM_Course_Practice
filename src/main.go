package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main() {
	db, err := gorm.Open("postgres", "user=postgres password=2wsx#EDC dbname=postgres sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	db.DropTable(&User{})
	db.CreateTable(&User{})

	db.Create(&User{
		FirstName: "Tricia",
		LastName:  "MacMillan-Dent",
	})

	db.Create(&User{
		FirstName: "Arthur",
		LastName:  "MacMillan-Dent",
	})

	db.Debug().Where("last_name LIKE ?", "Mac%").Delete(&User{})

}

type User struct {
	gorm.Model
	FirstName string
	LastName  string
}
