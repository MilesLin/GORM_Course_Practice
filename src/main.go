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

	u := User{
		FirstName: "Marvin",
		LastName:  "Robot",
	}

	tx := db.Begin()
	if err = tx.Debug().Create(&u).Error; err != nil {
		tx.Rollback()
	}

	u.LastName = "The Happy Robot"

	if err = tx.Debug().Save(&u).Error; err != nil {
		tx.Rollback()
	}

	tx.Commit()

}

type User struct {
	ID        uint
	FirstName string
	LastName  string
}
