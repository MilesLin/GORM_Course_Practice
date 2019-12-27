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

	u := User{
		FirstName: "Arthur",
		LastName:  "Dent",
	}

	db.Create(&u)

	// 如果已經新增了，就不會再度新增，且會回傳 false
	fmt.Println(db.NewRecord(&u))
}

type User struct {
	gorm.Model
	FirstName string
	LastName  string
}
