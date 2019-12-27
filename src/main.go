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
	db.DropTable(&Calendar{})
	db.CreateTable(&Calendar{})
	db.DropTable(&Appointment{})
	db.CreateTable(&Appointment{})
	db.Model(&Appointment{}).
		AddForeignKey("calendar_id", "calendars(id)", "RESTRICT", "RESTRICT")

	db.Debug().Save(&User{
		Username: "adent",
		Calendar: Calendar{
			Name: "Improbable Events",
			Appointments: []Appointment{
				{Subject: "Spontaneous Whale Generation"},
				{Subject: "Saved from Vaccuum of Space"},
			},
		},
	})

}

type User struct {
	gorm.Model
	Username  string
	FirstName string
	LastName  string
	Calendar  Calendar
}

type Calendar struct {
	gorm.Model
	Name         string
	UserID       uint
	Appointments []Appointment
}

type Appointment struct {
	gorm.Model
	Subject     string
	Description string
	StartTime   time.Time
	Length      uint
	CalendarID  uint
}
