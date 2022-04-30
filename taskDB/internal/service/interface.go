package service

import "taskDB/internal/model"

type Service interface {
	GetAll() ([]*model.BusFlight, error)
	GetByDate(flight *model.BusFlight) ([]*model.BusFlight, error)
	GetInformation(flight *model.BusFlight) ([]*model.BusFlight, error)
}
