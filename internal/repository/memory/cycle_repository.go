package memory

import (
	"context"
	"fmt"
	"sync"

	"git.home/alex/go-subscriptions/internal/domain/cycle/entity"
	"git.home/alex/go-subscriptions/internal/domain/cycle/repository"
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

func (r *CycleRepository) Create(ctx context.Context, cycle entity.Cycle) (*entity.Cycle, error) {
	r.Lock()
	defer r.Unlock()

	cycle.ID = uint(len(r.cycles) + 1)
	r.cycles[cycle.ID] = cycle

	return &cycle, nil
}

func (r *CycleRepository) Get(ctx context.Context, ID uint) (*entity.Cycle, error) {
	r.Lock()
	defer r.Unlock()

	cycle, ok := r.cycles[ID]
	if !ok {
		return nil, fmt.Errorf("cycle not found: %w", repository.ErrNotFoundCycle)
	}

	return &cycle, nil
}

func (r *CycleRepository) GetAll(ctx context.Context) (repository.Cycles, error) {
	r.Lock()
	defer r.Unlock()

	var cycles repository.Cycles
	for _, cycle := range r.cycles {
		cycles = append(cycles, cycle)
	}

	return cycles, nil
}

func (r *CycleRepository) Update(ctx context.Context, cycle entity.Cycle) (*entity.Cycle, error) {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.cycles[cycle.ID]; !ok {
		return nil, fmt.Errorf("cycle not found: %w", repository.ErrUpdateCycle)
	}

	r.cycles[cycle.ID] = cycle

	return &cycle, nil
}

func (r *CycleRepository) Delete(ctx context.Context, ID uint) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.cycles[ID]; !ok {
		return fmt.Errorf("cycle not found: %w", repository.ErrDeleteCycle)
	}

	delete(r.cycles, ID)

	return nil
}
