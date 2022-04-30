package database

import (
	"database/sql"
	"log"

	"github.com/Zhastlek/school21/internal/model"
)

type Storage interface {
	GetAllFlights() ([]*model.BusFlight, error)
	GetFlightsByDate(flight *model.BusFlight) ([]*model.BusFlight, error)
	GetInformationFlight(flight *model.BusFlight) ([]*model.BusFlight, error)
}

type storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) Storage {
	return &storage{
		db: db,
	}
}

func (s *storage) GetAllFlights() ([]*model.BusFlight, error) {
	rows, err := s.db.Query(`SELECT bus_flights.id, bus_flights.departure_city, bus_flights.arrival_city, bus_flights.distance,
        bus_flights.travel_time, time_flights.start_flight
		FROM bus_flights
			LEFT JOIN time_flights on bus_flights.id = time_flights.id
		ORDER BY bus_flights.id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	bf := []*model.BusFlight{}
	for rows.Next() {
		oneBf := &model.BusFlight{}
		err = rows.Scan(&oneBf.Id, &oneBf.DepartureCity, &oneBf.ArrivalCity, &oneBf.Distance, &oneBf.TravelTime, &oneBf.StartTrip)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Printf("ERROR storage GetAllFlights method :--> %v\n", err)
				return nil, err
			}
			return nil, err
		}
		bf = append(bf, oneBf)
	}
	return bf, nil
}

func (s *storage) GetFlightsByDate(flight *model.BusFlight) ([]*model.BusFlight, error) {
	rows, err := s.db.Query(`SELECT bus_flights.id, bus_flights.departure_city, bus_flights.arrival_city, bus_flights.distance,
       bus_flights.travel_time, COUNT(intermediate_bus_station.id), time_flights.start_flight
		FROM
			time_flights LEFT JOIN bus_flights on time_flights.id = bus_flights.id
						 LEFT JOIN intermediate_bus_station on bus_flights.id = intermediate_bus_station.id
		GROUP BY bus_flights.id
		HAVING bus_flights.departure_city = $1 and bus_flights.arrival_city = $2 and time_flights.start_flight = $3
		ORDER BY bus_flights.id;`, flight.DepartureCity, flight.ArrivalCity, flight.StartTrip)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	bf := []*model.BusFlight{}
	for rows.Next() {
		oneBf := &model.BusFlight{}

		err = rows.Scan(&oneBf.Id, &oneBf.DepartureCity, &oneBf.ArrivalCity, &oneBf.Distance, &oneBf.TravelTime, &oneBf.SumStation, &oneBf.StartTrip)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Printf("ERROR storage GetAllFlights method :--> %v\n", err)
				return nil, err
			}
			return nil, err
		}
		bf = append(bf, oneBf)
	}
	return bf, nil
}

func (s *storage) GetInformationFlight(flight *model.BusFlight) ([]*model.BusFlight, error) {
	rows, err := s.db.Query(`SELECT bus_flights.id, bus_flights.departure_city, bus_flights.arrival_city,
       bus_flights.distance, bus_flights.travel_time, COUNT(intermediate_bus_station.id), time_flights.start_flight
		FROM
			time_flights LEFT JOIN bus_flights on time_flights.id = bus_flights.id
						 LEFT JOIN intermediate_bus_station on bus_flights.id = intermediate_bus_station.id
		GROUP BY bus_flights.id
		HAVING bus_flights.departure_city = $1 and bus_flights.arrival_city = $2 and time_flights.start_flight= $3
		ORDER BY bus_flights.id`, flight.DepartureCity, flight.ArrivalCity, flight.StartTrip)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	bf := []*model.BusFlight{}
	for rows.Next() {
		oneBf := &model.BusFlight{}
		err = rows.Scan(&oneBf.Id, &oneBf.DepartureCity, &oneBf.ArrivalCity, &oneBf.Distance, &oneBf.TravelTime, &oneBf.SumStation, &oneBf.StartTrip)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Printf("ERROR storage GetAllFlights method :--> %v\n", err)
				return nil, err
			}
			return nil, err
		}
		stations, err := s.getIntermediateStations(oneBf.Id)
		if err != nil {
			log.Printf("ERROR storage GetAllFlights method getIntermediate stations:--> %v\n", err)
		}
		oneBf.Stations = append(oneBf.Stations, stations...)
		bf = append(bf, oneBf)
	}
	return bf, nil
}

func (s *storage) getIntermediateStations(id int) ([]model.IntermediateStation, error) {
	rows, err := s.db.Query(`SELECT station, number_station, arrival_time, departure_time
		FROM intermediate_bus_station 
		WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	stations := []model.IntermediateStation{}
	for rows.Next() {
		station := model.IntermediateStation{}
		err = rows.Scan(&station.NameStation, &station.Number, &station.ArrivalTime, &station.DepartureTime)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Printf("ERROR storage GetAllFlights method :--> %v\n", err)
				return stations, nil
			}
			return nil, err
		}
		stations = append(stations, station)
	}
	return stations, nil
}
