package repository

import (
	"context"

	"git.home/alex/go-subscriptions/internal/domain/cycle/entity"
)

type Cycles []entity.Cycle

type CycleRepository interface {
	Create(ctx context.Context, cycle entity.Cycle) (*entity.Cycle, error)
	GetByID(ctx context.Context, ID uint) (*entity.Cycle, error)
	GetAll(ctx context.Context) (Cycles, error)
	Update(ctx context.Context, cycle entity.Cycle) (*entity.Cycle, error)
	Delete(ctx context.Context, ID uint) error
}
