package repository

import (
	"context"
	"errors"

	"git.home/alex/go-subscriptions/internal/domain/entity"
)

var (
	ErrNotFoundCycle = errors.New("the cycle was not found in the repository")
	ErrCreateCycle   = errors.New("failed to add the cycle to the repository")
	ErrUpdateCycle   = errors.New("failed to update the cycle in the repository")
	ErrDeleteCycle   = errors.New("failed to delete the cycle from the repository")
)

type Cycles []entity.Cycle

type CycleRepository interface {
	Create(ctx context.Context, cycle entity.Cycle) (*entity.Cycle, error)
	Get(ctx context.Context, ID uint) (*entity.Cycle, error)
	GetAll(ctx context.Context) (Cycles, error)
	Update(ctx context.Context, cycle entity.Cycle) (*entity.Cycle, error)
	Delete(ctx context.Context, ID uint) error
}
