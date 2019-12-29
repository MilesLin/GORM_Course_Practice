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

	// appointment_user 是指多對多的那個 table
	// db.DropTableIfExists(&User{}, &Calendar{}, &Appointment{}, "appointment_user")
	db.DropTableIfExists("attachments")
	db.AutoMigrate(&User{}, &Calendar{}, &Appointment{})
	db.Debug().Model(&User{}).ModifyColumn("first_name", "VARCHAR(100)")

}

type User struct {
	gorm.Model
	Username  string
	FirstName string `sql:"type:VARCHAR(100)"`
	LastName  string
	Calendar  Calendar
}

type Calendar struct {
	gorm.Model
	Name         string
	UserID       uint `sql:"index:idx_calendar_user_id"`
	Appointments []*Appointment
}

type Appointment struct {
	gorm.Model
	Subject           string
	Description       string
	StartTime         time.Time
	Length            uint
	CalendarID        uint `sql:"index:idx_appointment_calendar_id"`
	Recurring         bool
	RecurrencePattern string
	Attendees         []*User `gorm:"many2many:appointment_user"`
}
