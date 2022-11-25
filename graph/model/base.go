package model

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB //database

func init() {

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=disable password=%s", dbHost, username, dbName, dbPort, password) //Build connection string

	log.Println(dbUri)
	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		log.Println(err)
	}

	db = conn
	err = db.Exec(`CREATE TABLE IF NOT EXISTS users(
		id serial PRIMARY KEY,
		username VARCHAR(50) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL
		);`).Error
	if err != nil {
		log.Fatal(err)
	}
	err = db.Exec(`CREATE TABLE IF NOT EXISTS todos(
		id serial PRIMARY KEY,
		title VARCHAR(25) NOT NULL,
		description VARCHAR(255),
		completed BOOLEAN NOT NULL,
		user_id int NOT NULL,
		CONSTRAINT fk_user
		FOREIGN KEY(user_id)
		REFERENCES users(id)
		ON DELETE CASCADE
		);`).Error
	if err != nil {
		log.Fatal(err)
	}
}

//returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
