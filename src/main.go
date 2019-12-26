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

	// db.SingularTable(true)

	db.CreateTable(&User{})

	for _, user := range users {
		db.Create(&user)
	}

	db.Where(&User{Username: "tmacmillan"}).Delete(&User{})

}

type User struct {
	gorm.Model
	Username  string `sql:"type:VARCHAR(15)"`
	FirstName string `sql:"size:100"`
	LastName  string
	Count     int  `gorm:"AUTO_INCREMENT"`
	TempField bool `sql:"-"`
}

func (u User) TableName() string {
	return "stackeholders"
}

var users []User = []User{
	User{Username: "adent", FirstName: "Arthur", LastName: "Dent"},
	User{Username: "fprefect", FirstName: "Ford", LastName: "Prefect"},
	User{Username: "tmacmillan", FirstName: "Tricia", LastName: "MacMillan"},
	User{Username: "mrobot", FirstName: "Marvin", LastName: "Robot"},
}
