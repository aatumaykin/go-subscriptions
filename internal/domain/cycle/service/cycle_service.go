package service

import (
	"context"

	"git.home/alex/go-subscriptions/internal/domain/cycle/entity"
	"git.home/alex/go-subscriptions/internal/domain/cycle/repository"
)

type CycleService interface {
	Create(ctx context.Context, cycle entity.Cycle) (*entity.Cycle, error)
	Get(ctx context.Context, ID uint) (*entity.Cycle, error)
	GetAll(ctx context.Context) (repository.Cycles, error)
	Update(ctx context.Context, cycle entity.Cycle) (*entity.Cycle, error)
	Delete(ctx context.Context, ID uint) error
}
