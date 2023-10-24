package db

import "database/sql"


func ConnectDatabase() *sql.DB {
	connection := "user=admin dbname=go-shop password=root port=5433 sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}
	return db
}