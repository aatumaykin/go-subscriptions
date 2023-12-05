package service

import (
	"context"

	"git.home/alex/go-subscriptions/internal/domain/category/repository"
)

type CollectionGetter interface {
	GetAll(ctx context.Context) (repository.Categories, error)
}
