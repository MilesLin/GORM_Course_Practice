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
	db.CreateTable(&User{})

	println("done")
}

type User struct {
	ID        uint
	Username  string
	FirstName string
	LastName  string
}
