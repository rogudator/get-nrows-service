package service

import "github.com/rogudator/get-nrows-service/internal/repository"

type Service struct {
	RowsGetter
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}
