package handlers

import (
	"net/http"

	"github.com/Zhastlek/school21/internal/service"
)

type Handler interface {
	Register(router *http.ServeMux)
}

func NewHandler(service service.FlightServiceInterface) Handler {
	return &handler{
		service: service,
	}
}

type handler struct {
	service service.FlightServiceInterface
}

func (h *handler) Register(router *http.ServeMux) {
	router.HandleFunc("/", h.HomePage)
	router.HandleFunc("/information", h.GetFlights)
}
