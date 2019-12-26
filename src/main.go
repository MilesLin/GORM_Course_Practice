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
	// db.DropTable(&User{})
	//
	// db.CreateTable(&User{})
	//
	// db.Model(&User{}).AddIndex("idx_first_name", "first_name")
	// db.Model(&User{}).AddUniqueIndex("idx_last_name", "last_name")

	db.Model(&User{}).RemoveIndex("idx_first_name")
}

type User struct {
	Model     gorm.Model
	FirstName string
	LastName  string
}

func (u User) TableName() string {
	return "stackeholders"
}

//
// var users []User = []User{
// 	User{Username: "adent", FirstName: "Arthur", LastName: "Dent"},
// 	User{Username: "fprefect", FirstName: "Ford", LastName: "Prefect"},
// 	User{Username: "tmacmillan", FirstName: "Tricia", LastName: "MacMillan"},
// 	User{Username: "mrobot", FirstName: "Marvin", LastName: "Robot"},
// }
