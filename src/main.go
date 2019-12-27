package main

import (
	"fmt"
	"time"

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
	db.DropTable(&Appointment{})
	db.CreateTable(&Appointment{})

	u := User{
		FirstName: "Ford",
		LastName:  "Prefect",
	}

	db.Create(&u)

	fmt.Println(u)
	fmt.Println()

	// db.Debug().Model(&u).Update("first_name", "Zaphod")
	db.Debug().Model(&u).Updates(
		map[string]interface{}{
			"first_name": "Zaphod",
			"last_name":  "Beeblebrox",
		})

	fmt.Println(u)
}

type User struct {
	gorm.Model
	FirstName    string
	LastName     string
	Appointments []Appointment
}

type Appointment struct {
	gorm.Model
	UserID      uint
	StartTime   *time.Time
	Duration    uint
	Attendees   []*User
	Subject     string
	Description string
}
