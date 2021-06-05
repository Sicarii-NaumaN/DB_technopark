package application

import (
	"forum/domain/entity"
	"forum/domain/repository"
)

type ServiceApp struct {
	s repository.ServiceRepository
}

func NewServiceApp(s repository.ServiceRepository) *ServiceApp {
	return &ServiceApp{s: s}
}

type ServiceAppInterface interface {
	ClearAllDate() error
	GetDBStatus() (*entity.Status, error)
}

func (s *ServiceApp) ClearAllDate() error {
	return nil
}

func (s *ServiceApp) GetDBStatus() (*entity.Status, error) {
	return nil, nil
}
