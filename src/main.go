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

	u := &User{
		FirstName: "Perfect",
		LastName:  "Ford",
	}
	db.Create(&u)
	db.Debug().Delete(&u)

}

type User struct {
	gorm.Model
	FirstName string
	LastName  string
}
