package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite3"
)

// User model for ORM
type User struct {
	gorm.Model
	Roles     string `json:roles`
	FirstName string `json: first_name`
	LastName  string `json: last_name`
	Email     string `json: email`
	password  string
	BirthDate string `json:birth_date`
}

func initMigration() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect to db")
	}
	// close database
	defer db.Close()

	//create table
	db.AutoMigrate(User{})
}
