package memory

import (
	"context"
	"sync"

	"git.home/alex/go-subscriptions/internal/domain/entity"
	"git.home/alex/go-subscriptions/internal/domain/repository"
)

type CycleRepository struct {
	cycles map[uint]entity.Cycle
	sync.Mutex
}

func NewCycleRepository() *CycleRepository {
	return &CycleRepository{
		cycles: make(map[uint]entity.Cycle),
	}
}

func (r *CycleRepository) Create(_ context.Context, cycle entity.Cycle) (*entity.Cycle, error) {
	r.Lock()
	defer r.Unlock()

	cycle.ID = uint(len(r.cycles) + 1)
	r.cycles[cycle.ID] = cycle

	return &cycle, nil
}

func (r *CycleRepository) Get(_ context.Context, id uint) (*entity.Cycle, error) {
	r.Lock()
	defer r.Unlock()

	cycle, ok := r.cycles[id]
	if !ok {
		return nil, repository.ErrNotFoundCycle
	}

	return &cycle, nil
}

func (r *CycleRepository) GetAll(_ context.Context) (repository.Cycles, error) {
	r.Lock()
	defer r.Unlock()

	var cycles repository.Cycles
	for _, cycle := range r.cycles {
		cycles = append(cycles, cycle)
	}

	return cycles, nil
}

func (r *CycleRepository) Update(_ context.Context, cycle entity.Cycle) (*entity.Cycle, error) {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.cycles[cycle.ID]; !ok {
		return nil, repository.ErrNotFoundCycle
	}

	r.cycles[cycle.ID] = cycle

	return &cycle, nil
}

func (r *CycleRepository) Delete(_ context.Context, id uint) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.cycles[id]; !ok {
		return repository.ErrNotFoundCycle
	}

	delete(r.cycles, id)

	return nil
}
