package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/Zhastlek/school21/internal/model"
)

func (h *handler) GetFlights(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	log.Println(r.Method, r.FormValue("information"), "|", r.FormValue("trip-start"), "|")

	if r.FormValue("information") == "information" {
		log.Println("case information")
		h.GetInformationFlight(w, r)
	} else if r.FormValue("trip-start") != "" {
		log.Println("case trip-start")
		h.GetFlightsByDate(w, r)
	} else {
		log.Println("case all")
		h.GetAllFlights(w, r)
	}
}

func (h *handler) GetAllFlights(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	flights, err := h.service.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res := &model.Result{
		Flights:         flights,
		IsExistFlight:   true,
		PossibleFlights: false,
		FullInformation: false,
	}
	temp.Execute(w, res)
}

func (h *handler) GetFlightsByDate(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	departureCity := r.FormValue("departure_city")
	arrivalCity := r.FormValue("arrival_city")
	time := r.FormValue("trip-start")

	// log.Println(departureCity, arrivalCity, "handler")
	bf := &model.BusFlight{
		DepartureCity: departureCity,
		ArrivalCity:   arrivalCity,
		StartTrip:     time,
	}

	flights, err := h.service.GetByDate(bf)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res := &model.Result{
		Flights:         flights,
		IsExistFlight:   true,
		PossibleFlights: true,
		FullInformation: false,
	}
	temp.Execute(w, res)
}

func (h *handler) GetInformationFlight(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	departureCity := r.FormValue("departure_city")
	arrivalCity := r.FormValue("arrival_city")
	time := r.FormValue("trip-start")

	bf := &model.BusFlight{
		DepartureCity: departureCity,
		ArrivalCity:   arrivalCity,
		StartTrip:     time,
	}

	flights, err := h.service.GetInformation(bf)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res := &model.Result{
		Flights:         flights,
		IsExistFlight:   true,
		PossibleFlights: true,
		FullInformation: true,
	}

	temp.Execute(w, res)
}

func (h *handler) HomePage(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	res := &model.Result{}
	temp.Execute(w, res)
}
