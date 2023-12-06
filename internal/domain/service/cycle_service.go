package service

import (
	"context"
	"errors"

	"git.home/alex/go-subscriptions/internal/domain/entity"
	"git.home/alex/go-subscriptions/internal/domain/repository"
)

var (
	ErrInvalidCycle = errors.New("the cycle is invalid")
)

type CycleService struct {
	repo repository.CycleRepository
}

func NewCycleService(repo repository.CycleRepository) *CycleService {
	return &CycleService{repo: repo}
}

func (s *CycleService) CreateCycle(ctx context.Context, cycle entity.Cycle) (*entity.Cycle, error) {
	if cycle.Name == "" {
		return nil, ErrInvalidCycle
	}

	return s.repo.Create(ctx, cycle)
}

func (s *CycleService) GetCycle(ctx context.Context, id uint) (*entity.Cycle, error) {
	if id == 0 {
		return nil, repository.ErrNotFoundCycle
	}

	return s.repo.Get(ctx, id)
}

func (s *CycleService) GetAllCycles(ctx context.Context) (repository.Cycles, error) {
	return s.repo.GetAll(ctx)
}

func (s *CycleService) UpdateCycle(ctx context.Context, cycle entity.Cycle) (*entity.Cycle, error) {
	if cycle.ID == 0 || cycle.Name == "" {
		return nil, ErrInvalidCycle
	}

	return s.repo.Update(ctx, cycle)
}

func (s *CycleService) DeleteCycle(ctx context.Context, id uint) error {
	if id == 0 {
		return repository.ErrNotFoundCycle
	}

	return s.repo.Delete(ctx, id)
}
