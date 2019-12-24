package main

import (
	"fmt"

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

	user := User{
		Username:  "adent",
		FirstName: "Arthur",
		LastName:  "Dent",
	}

	fmt.Println(user)

	db.Create(&user)

	fmt.Println(user)

	println("done")
}

type User struct {
	ID        uint
	Username  string
	FirstName string
	LastName  string
}
