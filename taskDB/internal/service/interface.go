package service

import "github.com/Zhastlek/school21/internal/model"

type FlightServiceInterface interface {
	GetAll() ([]*model.BusFlight, error)
	GetByDate(flight *model.BusFlight) ([]*model.BusFlight, error)
	GetInformation(flight *model.BusFlight) ([]*model.BusFlight, error)
}
