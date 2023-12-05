package service

import (
	"context"

	"git.home/alex/go-subscriptions/internal/domain/category/entity"
)

type Creator interface {
	Create(ctx context.Context, name string) (*entity.Category, error)
}
