package handlers

import (
	"html/template"
	"log"
	"net/http"
	"taskDB/internal/model"
	"taskDB/internal/service"
)

type Handler interface {
	Register(router *http.ServeMux)
}

type handler struct {
	service service.Service
}
type Data struct {
	Flights         []*model.BusFlight
	Flight          bool
	PossibleFlights bool
	FullInformation bool
}

func NewHandler(service service.Service) Handler {
	return &handler{
		service: service,
	}
}

func (h *handler) Register(router *http.ServeMux) {
	router.HandleFunc("/", h.HomePage)
	router.HandleFunc("/information", h.Get)
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		return
	}
	log.Println(r.Method, r.FormValue("information"), "|", r.FormValue("trip-start"), "|")
	if r.FormValue("information") == "information" {
		log.Println("information")
		h.GetInformationFlight(w, r)
	} else if r.FormValue("trip-start") != "" {
		log.Println("trip-start")
		h.GetFlightsByDate(w, r)
	} else {
		log.Println("all")
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
	//log.Println(flights[0])
	data := &Data{
		flights,
		true,
		false,
		false,
	}
	temp.Execute(w, data)
}

func (h *handler) GetFlightsByDate(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	departureCity := r.FormValue("departure_city")
	arrivalCity := r.FormValue("arrival_city")
	log.Println(departureCity, arrivalCity, "handler")
	time := r.FormValue("trip-start")
	bf := &model.BusFlight{}
	bf.DepartureCity = departureCity
	bf.ArrivalCity = arrivalCity
	bf.StartTrip = time
	flights, err := h.service.GetByDate(bf)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	data := &Data{
		flights,
		true,
		true,
		false,
	}
	temp.Execute(w, data)
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
	bf := &model.BusFlight{}
	bf.DepartureCity = departureCity
	bf.ArrivalCity = arrivalCity
	bf.StartTrip = time
	flights, err := h.service.GetInformation(bf)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	data := &Data{
		flights,
		true,
		true,
		true,
	}
	temp.Execute(w, data)
}

func (h *handler) HomePage(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data := &Data{}
	temp.Execute(w, data)
}
