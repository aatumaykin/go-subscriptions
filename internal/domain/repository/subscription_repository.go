package repository

import (
	"context"
	"errors"

	"git.home/alex/go-subscriptions/internal/domain/entity"
)

var (
	ErrNotFoundSubscription = errors.New("the subscription was not found in the repository")
	ErrCreateSubscription   = errors.New("failed to add the subscription to the repository")
	ErrUpdateSubscription   = errors.New("failed to update the subscription in the repository")
	ErrDeleteSubscription   = errors.New("failed to delete the subscription from the repository")
)

type Subscriptions []entity.Subscription

type SubscriptionRepository interface {
	Create(ctx context.Context, subscription entity.Subscription) (*entity.Subscription, error)
	Get(ctx context.Context, ID uint) (*entity.Subscription, error)
	GetAll(ctx context.Context) (Subscriptions, error)
	Update(ctx context.Context, subscription entity.Subscription) (*entity.Subscription, error)
	Delete(ctx context.Context, ID uint) error
}
