package taskdb

import (
	"database/sql"
	"net/http"
	"taskDB/internal/adapters/database"
	"taskDB/internal/adapters/handlers"
	"taskDB/internal/service"
)

const speed = 120

func Config(db *sql.DB) *http.ServeMux {
	router := http.NewServeMux()
	storage := database.NewStorage(db)
	newService := service.NewService(storage)
	handler := handlers.NewHandler(newService)
	handler.Register(router)
	return router
}
