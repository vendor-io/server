package db

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init() {
	db, err := gorm.Open(sqlite.Open("keyboardify_gorm.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}

	var schema = `
		  DROP TABLE IF EXISTS users;
		  CREATE TABLE users (
			 ID                  INTEGER PRIMARY KEY
		 	);
		
		  DROP TABLE IF EXISTS keyboards;
		  CREATE TABLE keyboards (
			 ID                  INTEGER PRIMARY KEY
			);
	`

	db.Exec(schema)
}
