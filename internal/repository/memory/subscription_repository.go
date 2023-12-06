package memory

import (
	"context"
	"sync"

	"git.home/alex/go-subscriptions/internal/domain/entity"
	"git.home/alex/go-subscriptions/internal/domain/repository"
)

type SubscriptionRepository struct {
	subscriptions map[uint]entity.Subscription
	sync.Mutex
}

func NewSubscriptionRepository() *SubscriptionRepository {
	return &SubscriptionRepository{
		subscriptions: make(map[uint]entity.Subscription),
	}
}

func (r *SubscriptionRepository) Create(_ context.Context, subscription entity.Subscription) (*entity.Subscription, error) {
	r.Lock()
	defer r.Unlock()

	subscription.ID = uint(len(r.subscriptions) + 1)
	r.subscriptions[subscription.ID] = subscription

	return &subscription, nil
}

func (r *SubscriptionRepository) Get(_ context.Context, id uint) (*entity.Subscription, error) {
	r.Lock()
	defer r.Unlock()

	subscription, ok := r.subscriptions[id]
	if !ok {
		return nil, repository.ErrNotFoundSubscription
	}

	return &subscription, nil
}

func (r *SubscriptionRepository) GetAll(_ context.Context) (repository.Subscriptions, error) {
	r.Lock()
	defer r.Unlock()

	var subscriptions repository.Subscriptions
	for _, subscription := range r.subscriptions {
		subscriptions = append(subscriptions, subscription)
	}

	return subscriptions, nil
}

func (r *SubscriptionRepository) Update(_ context.Context, subscription entity.Subscription) (*entity.Subscription, error) {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.subscriptions[subscription.ID]; !ok {
		return nil, repository.ErrUpdateSubscription
	}

	r.subscriptions[subscription.ID] = subscription

	return &subscription, nil
}

func (r *SubscriptionRepository) Delete(_ context.Context, id uint) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.subscriptions[id]; !ok {
		return repository.ErrDeleteSubscription
	}

	delete(r.subscriptions, id)

	return nil
}
