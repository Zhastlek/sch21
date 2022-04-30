package service

import (
	"taskDB/internal/adapters/database"
	"taskDB/internal/model"
)

type service struct {
	db database.Storage
}

func NewService(storage database.Storage) Service {
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
