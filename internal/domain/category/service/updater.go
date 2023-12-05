package service

import (
	"context"

	"git.home/alex/go-subscriptions/internal/domain/category/entity"
)

type Updater interface {
	Update(ctx context.Context, category entity.Category) (*entity.Category, error)
}
