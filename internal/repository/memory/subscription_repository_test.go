package memory_test

import (
	"context"
	"testing"

	"git.home/alex/go-subscriptions/internal/domain/entity"
	"git.home/alex/go-subscriptions/internal/domain/repository"
	"git.home/alex/go-subscriptions/internal/repository/memory"
	"github.com/stretchr/testify/assert"
)

func TestSubscriptionRepository_Create(t *testing.T) {
	testCases := []struct {
		name         string
		subscription entity.Subscription
		wantResult   *entity.Subscription
		wantErr      error
	}{
		{
			name:         "Create a new subscription",
			subscription: entity.Subscription{Name: "Test Subscription"},
			wantResult:   &entity.Subscription{ID: 1, Name: "Test Subscription"},
			wantErr:      nil,
		},
		{
			name:         "Create a new subscription",
			subscription: entity.Subscription{Name: "Test Subscription"},
			wantResult:   &entity.Subscription{ID: 2, Name: "Test Subscription"},
			wantErr:      nil,
		},
	}

	repo := memory.NewSubscriptionRepository()
	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := repo.Create(ctx, tc.subscription)

			if tc.wantErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.wantErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.wantResult, result)
			}
		})
	}
}

func TestSubscriptionRepository_Get(t *testing.T) {
	testCases := []struct {
		name         string
		subscription entity.Subscription
		id           uint
		wantErr      error
	}{
		{
			name:         "Get an existing subscription",
			subscription: entity.Subscription{Name: "Test Subscription"},
			id:           1,
			wantErr:      nil,
		},
		{
			name:         "Get an existing subscription",
			subscription: entity.Subscription{Name: "Test Subscription"},
			id:           2,
			wantErr:      nil,
		},
		{
			name:         "Get a non-existing subscription",
			subscription: entity.Subscription{Name: "Test Subscription"},
			id:           10,
			wantErr:      repository.ErrNotFoundSubscription,
		},
	}

	repo := memory.NewSubscriptionRepository()
	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			createdSubs, err := repo.Create(ctx, tc.subscription)
			assert.NoError(t, err)

			foundSubs, err := repo.Get(ctx, tc.id)

			if tc.wantErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.wantErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, createdSubs, foundSubs)
			}
		})
	}
}

func TestSubscriptionRepository_GetAll(t *testing.T) {
	testCases := []struct {
		name          string
		subscriptions repository.Subscriptions
		wantResult    repository.Subscriptions
		expectedLen   int
	}{
		{
			name:        "Empty repository",
			expectedLen: 0,
		},
		{
			name:          "Get all subscriptions",
			subscriptions: []entity.Subscription{{Name: "Subscription 1"}, {Name: "Subscription 2"}},
			wantResult:    []entity.Subscription{{ID: 1, Name: "Subscription 1"}, {ID: 2, Name: "Subscription 2"}},
			expectedLen:   2,
		},
	}

	repo := memory.NewSubscriptionRepository()
	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for _, subscription := range tc.subscriptions {
				_, err := repo.Create(ctx, subscription)
				assert.NoError(t, err)
			}

			subscriptions, err := repo.GetAll(ctx)
			assert.NoError(t, err)
			assert.Equal(t, tc.wantResult, subscriptions)
			assert.Equal(t, tc.expectedLen, len(subscriptions))
		})
	}
}

func TestSubscriptionRepository_Update(t *testing.T) {
	testCases := []struct {
		name                string
		initialSubscription entity.Subscription
		updatedSubscription entity.Subscription
		wantResult          *entity.Subscription
		wantErr             error
	}{
		{
			name:                "Update an existing subscription",
			initialSubscription: entity.Subscription{Name: "Test Subscription"},
			updatedSubscription: entity.Subscription{ID: 1, Name: "Updated Test Subscription"},
			wantResult:          &entity.Subscription{ID: 1, Name: "Updated Test Subscription"},
			wantErr:             nil,
		},
		{
			name:                "Update an existing subscription",
			initialSubscription: entity.Subscription{Name: "Test Subscription"},
			updatedSubscription: entity.Subscription{ID: 2, Name: "Updated Test Subscription"},
			wantResult:          &entity.Subscription{ID: 2, Name: "Updated Test Subscription"},
			wantErr:             nil,
		},
		{
			name:                "Update a non-existing subscription",
			initialSubscription: entity.Subscription{Name: "Test Subscription"},
			updatedSubscription: entity.Subscription{ID: 10, Name: "Updated Test Subscription"},
			wantResult:          nil,
			wantErr:             repository.ErrUpdateSubscription,
		},
	}

	repo := memory.NewSubscriptionRepository()
	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.Create(ctx, tc.initialSubscription)
			assert.NoError(t, err)

			result, err := repo.Update(ctx, tc.updatedSubscription)

			if tc.wantErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.wantErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.wantResult, result)
			}
		})
	}
}

func TestSubscriptionRepository_Delete(t *testing.T) {
	testCases := []struct {
		name         string
		subscription entity.Subscription
		id           uint
		wantErr      error
	}{
		{
			name:         "Delete an existing subscription",
			subscription: entity.Subscription{Name: "Test Subscription"},
			id:           1,
			wantErr:      nil,
		},
		{
			name:         "Delete an existing subscription",
			subscription: entity.Subscription{Name: "Test Subscription"},
			id:           1,
			wantErr:      nil,
		},
		{
			name:         "Delete a non-existing subscription",
			subscription: entity.Subscription{Name: "Test Subscription"},
			id:           10,
			wantErr:      repository.ErrDeleteSubscription,
		},
	}

	repo := memory.NewSubscriptionRepository()
	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.Create(ctx, tc.subscription)
			assert.NoError(t, err)

			err = repo.Delete(ctx, tc.id)

			if tc.wantErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.wantErr, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
