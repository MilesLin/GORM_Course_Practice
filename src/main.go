package main

import (
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

	db.Create(&User{
		FirstName: "Tricia",
		LastName:  "Dent",
		Salary:    50000,
	})
	db.Create(&User{
		FirstName: "Authur",
		LastName:  "Dent",
		Salary:    30000,
	})

	db.Debug().Table("users").Where("last_name = ?", "Dent").
		Update("last_name", "MacMillan-Dent")

}

type User struct {
	gorm.Model
	FirstName    string
	LastName     string
	Salary       uint
	Appointments []Appointment
}

func (u *User) BeforeUpdate() error {
	println("Before Update")
	return nil
}
func (u *User) AfterUpdate() error {
	println("After Update")
	return nil
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
