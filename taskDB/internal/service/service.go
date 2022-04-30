package service

import (
	"github.com/Zhastlek/school21/internal/adapters/database"
	"github.com/Zhastlek/school21/internal/model"
)

type service struct {
	db database.Storage
}

func NewFlightService(storage database.Storage) FlightServiceInterface {
	return &service{
		db: storage,
	}
}

func (s *service) GetAll() ([]*model.BusFlight, error) {
	bfs, err := s.db.GetAllFlights()
	if err != nil {
		return nil, err
	}
	return bfs, nil
}

func (s *service) GetByDate(b *model.BusFlight) ([]*model.BusFlight, error) {
	bfs, err := s.db.GetFlightsByDate(b)
	if err != nil {
		return nil, err
	}
	return bfs, nil
}

func (s *service) GetInformation(b *model.BusFlight) ([]*model.BusFlight, error) {
	bfs, err := s.db.GetInformationFlight(b)
	if err != nil {
		return nil, err
	}
	return bfs, nil
}
