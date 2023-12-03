package subscription

import (
	"context"

	"git.home/alex/go-subscriptions/internal/domain/subscription/entity"
	"git.home/alex/go-subscriptions/internal/domain/subscription/repository"
	"git.home/alex/go-subscriptions/internal/repository/memory"
)

type Service struct {
	repository repository.SubscriptionRepository
}

type ServiceConfiguration func(cs *Service) error

func NewSubscriptionService(cfgs ...ServiceConfiguration) (*Service, error) {
	cs := &Service{}

	// Apply all Configurations passed in
	for _, cfg := range cfgs {
		err := cfg(cs)
		if err != nil {
			return nil, err
		}
	}

	return cs, nil
}

func WithSubscriptionRepository(r repository.SubscriptionRepository) ServiceConfiguration {
	return func(cs *Service) error {
		cs.repository = r
		return nil
	}
}

func WithMemorySubscriptionRepository() ServiceConfiguration {
	return WithSubscriptionRepository(memory.NewSubscriptionRepository())
}

func (cs *Service) Create(ctx context.Context, subscription entity.Subscription) (*entity.Subscription, error) {
	return cs.repository.Create(ctx, subscription)
}

func (cs *Service) Get(ctx context.Context, ID uint) (*entity.Subscription, error) {
	return cs.repository.Get(ctx, ID)
}

func (cs *Service) GetAll(ctx context.Context) (repository.Subscriptions, error) {
	return cs.repository.GetAll(ctx)
}

func (cs *Service) Update(ctx context.Context, subscription entity.Subscription) (*entity.Subscription, error) {
	return cs.repository.Update(ctx, subscription)
}

func (cs *Service) Delete(ctx context.Context, ID uint) error {
	return cs.repository.Delete(ctx, ID)
}
