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

	dbase := db.DB()
	defer dbase.Close()

	err = dbase.Ping()
	if err != nil {
		panic(err.Error())
	}

	println("Connection to database established")
	// db, err := sql.Open("postgres", "postgres://postgres:'2wsx#EDC'@localhost/lss?sslmode=disable")
	// db, err := sql.Open("postgres", "user=postgres password=2wsx#EDC dbname=lss sslmode=disable")
}
