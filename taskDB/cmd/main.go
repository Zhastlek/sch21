package main

import (
	"log"
	"net/http"
	taskdb "taskDB"
	"taskDB/internal/adapters/database"
)

func main() {
	//var db *sql.DB
	db := database.CheckDB()
	router := taskdb.Config(db)
	log.Println("port is 8000 listening...")
	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Printf("%v\n", err)
	}
}
