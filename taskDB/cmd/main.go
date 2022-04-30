package main

import (
	"log"
	"net/http"

	"github.com/Zhastlek/school21/app"
	"github.com/Zhastlek/school21/internal/adapters/database"
)

func main() {
	db, err := database.InitDb()
	if err != nil {
		log.Fatal(err)
	}

	router := app.Initialize(db)
	log.Println("Server run  is 8000 port...")
	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatal(err)
	}
}
