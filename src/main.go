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

	for _, user := range users {
		db.Create(&user)
	}

	u := User{Username: "tmacmillan"}
	db.Where(&u).First(&u)

	// u := User{}
	// db.Where("username = ?", "tmacmillan").First(&u)

	fmt.Println(u)

	u.LastName = "Beeblebrox"

	db.Save(&u)

	user := User{}
	db.Where(&u).First(&user)

	fmt.Println(user)

}

type User struct {
	ID        uint
	Username  string
	FirstName string
	LastName  string
}

var users []User = []User{
	User{Username: "adent", FirstName: "Arthur", LastName: "Dent"},
	User{Username: "fprefect", FirstName: "Ford", LastName: "Prefect"},
	User{Username: "tmacmillan", FirstName: "Tricia", LastName: "MacMillan"},
	User{Username: "mrobot", FirstName: "Marvin", LastName: "Robot"},
}
