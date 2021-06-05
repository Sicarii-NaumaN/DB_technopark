package persistence

import (
	"forum/domain/entity"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ServiceRepo struct {
	db *pgxpool.Pool
}

func NewServiceRepository(db *pgxpool.Pool) *ServiceRepo {
	return &ServiceRepo{db: db}
}

func (s *ServiceRepo) ClearAllDate() error {
	return nil
}

func (s *ServiceRepo) GetDBStatus() (*entity.Status, error) {
	return nil, nil
}


