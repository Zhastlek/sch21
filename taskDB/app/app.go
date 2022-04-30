package app

import (
	"database/sql"
	"net/http"

	"github.com/Zhastlek/school21/internal/adapters/database"
	"github.com/Zhastlek/school21/internal/adapters/handlers"
	"github.com/Zhastlek/school21/internal/service"
)

func Initialize(db *sql.DB) *http.ServeMux {
	router := http.NewServeMux()
	storage := database.NewStorage(db)
	newService := service.NewFlightService(storage)
	handler := handlers.NewHandler(newService)
	handler.Register(router)
	return router
}
