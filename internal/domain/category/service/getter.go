package service

import (
	"context"

	"git.home/alex/go-subscriptions/internal/domain/category/entity"
)

type Getter interface {
	Get(ctx context.Context, ID uint) (*entity.Category, error)
}
