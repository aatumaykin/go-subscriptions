package memory

import (
	"context"
	"fmt"
	"sync"

	"git.home/alex/go-subscriptions/internal/domain/subscription/entity"
	"git.home/alex/go-subscriptions/internal/domain/subscription/repository"
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

func (r *SubscriptionRepository) Create(ctx context.Context, subscription entity.Subscription) (*entity.Subscription, error) {
	r.Lock()
	defer r.Unlock()

	subscription.ID = uint(len(r.subscriptions) + 1)
	r.subscriptions[subscription.ID] = subscription

	return &subscription, nil
}

func (r *SubscriptionRepository) Get(ctx context.Context, ID uint) (*entity.Subscription, error) {
	r.Lock()
	defer r.Unlock()

	subscription, ok := r.subscriptions[ID]
	if !ok {
		return nil, fmt.Errorf("subscription not found: %w", repository.ErrNotFoundSubscription)
	}

	return &subscription, nil
}

func (r *SubscriptionRepository) GetAll(ctx context.Context) (repository.Subscriptions, error) {
	r.Lock()
	defer r.Unlock()

	var subscriptions repository.Subscriptions
	for _, subscription := range r.subscriptions {
		subscriptions = append(subscriptions, subscription)
	}

	return subscriptions, nil
}

func (r *SubscriptionRepository) Update(ctx context.Context, subscription entity.Subscription) (*entity.Subscription, error) {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.subscriptions[subscription.ID]; !ok {
		return nil, fmt.Errorf("subscription not found: %w", repository.ErrUpdateSubscription)
	}

	r.subscriptions[subscription.ID] = subscription

	return &subscription, nil
}

func (r *SubscriptionRepository) Delete(ctx context.Context, ID uint) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.subscriptions[ID]; !ok {
		return fmt.Errorf("subscription not found: %w", repository.ErrDeleteSubscription)
	}

	delete(r.subscriptions, ID)

	return nil
}
