package repository

import (
	"context"

	"git.home/alex/go-subscriptions/internal/domain/subscription/entity"
)

type Subscriptions []entity.Subscription

type SubscriptionRepository interface {
	Create(ctx context.Context, subscription entity.Subscription) (*entity.Subscription, error)
	Get(ctx context.Context, ID uint) (*entity.Subscription, error)
	GetAll(ctx context.Context) (Subscriptions, error)
	Update(ctx context.Context, subscription entity.Subscription) (*entity.Subscription, error)
	Delete(ctx context.Context, ID uint) error
}
