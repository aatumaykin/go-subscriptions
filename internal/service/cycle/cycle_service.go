package cycle

import (
	"context"

	"git.home/alex/go-subscriptions/internal/domain/cycle/entity"
	"git.home/alex/go-subscriptions/internal/domain/cycle/repository"
	"git.home/alex/go-subscriptions/internal/repository/memory"
)

type Service struct {
	repository repository.CycleRepository
}

type ServiceConfiguration func(cs *Service) error

func NewCycleService(cfgs ...ServiceConfiguration) (*Service, error) {
	cs := &Service{}

	// Apply all Configurations passed in
	for _, cfg := range cfgs {
		err := cfg(cs)
		if err != nil {
			return nil, err
		}
	}

	return cs, nil
}

func WithCycleRepository(r repository.CycleRepository) ServiceConfiguration {
	return func(cs *Service) error {
		cs.repository = r
		return nil
	}
}

func WithMemoryCycleRepository() ServiceConfiguration {
	return WithCycleRepository(memory.NewCycleRepository())
}

func (cs *Service) Create(ctx context.Context, cycle entity.Cycle) (*entity.Cycle, error) {
	return cs.repository.Create(ctx, cycle)
}

func (cs *Service) Get(ctx context.Context, ID uint) (*entity.Cycle, error) {
	return cs.repository.Get(ctx, ID)
}

func (cs *Service) GetAll(ctx context.Context) (repository.Cycles, error) {
	return cs.repository.GetAll(ctx)
}

func (cs *Service) Update(ctx context.Context, cycle entity.Cycle) (*entity.Cycle, error) {
	return cs.repository.Update(ctx, cycle)
}

func (cs *Service) Delete(ctx context.Context, ID uint) error {
	return cs.repository.Delete(ctx, ID)
}
