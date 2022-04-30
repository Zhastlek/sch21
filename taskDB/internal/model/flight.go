package model

type BusFlight struct {
	Id            int
	DepartureCity string
	ArrivalCity   string
	StartTrip     string
	EndTrip       string
	Distance      string
	TravelTime    string
	SumStation    int
	Stations      []IntermediateStation
}

type IntermediateStation struct {
	Id            int
	Number        int
	NameStation   string
	ArrivalTime   string
	DepartureTime string
}

type Result struct {
	Flights         []*BusFlight
	IsExistFlight   bool
	PossibleFlights bool
	FullInformation bool
}
