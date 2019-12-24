package model

import (
	"fmt"
	"database/sql"
	 "log"
	
	_ "github.com/lib/pq"
)

const (
	host="localhost"
	port = 5432
	user= "yuanyu"
	password = "123456"
	dbname="redpill_db"
)
var con *sql.DB

func Connect() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	con = db
	fmt.Println("Successfully connected")
	return db
}