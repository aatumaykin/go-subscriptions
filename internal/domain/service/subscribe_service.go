package service

import (
	"context"
	"errors"

	"git.home/alex/go-subscriptions/internal/domain/entity"
	"git.home/alex/go-subscriptions/internal/domain/repository"
)

var (
	ErrInvalidSubscription = errors.New("the subscription is invalid")
)

type SubscriptionService struct {
	repo repository.SubscriptionRepository
}

func NewSubscriptionService(repo repository.SubscriptionRepository) *SubscriptionService {
	return &SubscriptionService{repo: repo}
}

func (s *SubscriptionService) CreateSubscription(ctx context.Context, subscription entity.Subscription) (*entity.Subscription, error) {
	if subscription.Price <= 0 || subscription.Name == "" {
		return nil, ErrInvalidSubscription
	}

	if subscription.Currency.Code == "" || subscription.Cycle.ID == 0 {
		return nil, ErrInvalidSubscription
	}

	return s.repo.Create(ctx, subscription)
}

func (s *SubscriptionService) GetSubscription(ctx context.Context, id uint) (*entity.Subscription, error) {
	return s.repo.Get(ctx, id)
}

func (s *SubscriptionService) GetAllSubscriptions(ctx context.Context) (repository.Subscriptions, error) {
	return s.repo.GetAll(ctx)
}

func (s *SubscriptionService) UpdateSubscription(ctx context.Context, subscription entity.Subscription) (*entity.Subscription, error) {
	if subscription.ID == 0 || subscription.Price <= 0 || subscription.Name == "" {
		return nil, ErrInvalidSubscription
	}

	if subscription.Currency.Code == "" || subscription.Cycle.ID == 0 {
		return nil, ErrInvalidSubscription
	}

	return s.repo.Update(ctx, subscription)
}

func (s *SubscriptionService) DeleteSubscription(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
