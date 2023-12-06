package memory_test

import (
	"context"
	"errors"
	"testing"

	"git.home/alex/go-subscriptions/internal/domain/entity"
	"git.home/alex/go-subscriptions/internal/domain/repository"
	"git.home/alex/go-subscriptions/internal/repository/memory"
)

func TestSubscriptionRepository_Create(t *testing.T) {
	type testCase struct {
		test         string
		subscription entity.Subscription
		expectedErr  error
	}

	testCases := []testCase{
		{
			test:         "Create a new subscription",
			subscription: entity.Subscription{Name: "Test Subscription"},
			expectedErr:  nil,
		},
		{
			test:         "Create a new subscription",
			subscription: entity.Subscription{ID: 1, Name: "Test Subscription"},
			expectedErr:  nil,
		},
	}

	repo := memory.NewSubscriptionRepository()

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			createdSubscription, err := repo.Create(context.Background(), tc.subscription)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("Expected error: %v, got: %v", tc.expectedErr, err)
			}

			if createdSubscription != nil {
				found, err := repo.Get(context.Background(), createdSubscription.ID)
				if err != nil {
					t.Fatal(err)
				}
				if createdSubscription.Name != found.Name {
					t.Errorf("Expected %v, got %v", createdSubscription.Name, found.Name)
				}
			}
		})
	}
}

func TestSubscriptionRepository_Get(t *testing.T) {
	type testCase struct {
		test         string
		subscription entity.Subscription
		expectedID   uint
		expectedErr  error
	}

	testCases := []testCase{
		{
			test:         "Get an existing subscription",
			subscription: entity.Subscription{Name: "Test Subscription"},
			expectedID:   1,
			expectedErr:  nil,
		},
		{
			test:         "Get a non-existing subscription",
			subscription: entity.Subscription{Name: "Test Subscription"},
			expectedID:   3,
			expectedErr:  repository.ErrNotFoundSubscription,
		},
	}

	repo := memory.NewSubscriptionRepository()

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := repo.Create(context.Background(), tc.subscription)
			if err != nil {
				t.Fatal(err)
			}

			_, err = repo.Get(context.Background(), tc.expectedID)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("Expected error: %v, got: %v", tc.expectedErr, err)
			}
		})
	}
}

func TestSubscriptionRepository_GetAll(t *testing.T) {
	type testCase struct {
		test          string
		subscriptions []entity.Subscription
		expectedLen   int
	}

	testCases := []testCase{
		{
			test:        "Empty repository",
			expectedLen: 0,
		},
		{
			test:          "Get all subscriptions",
			subscriptions: []entity.Subscription{{Name: "Subscription 1"}, {Name: "Subscription 2"}},
			expectedLen:   2,
		},
	}

	repo := memory.NewSubscriptionRepository()

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			for _, subscription := range tc.subscriptions {
				_, err := repo.Create(context.Background(), subscription)
				if err != nil {
					t.Fatal(err)
				}
			}

			subscriptions, _ := repo.GetAll(context.Background())
			if len(subscriptions) != tc.expectedLen {
				t.Errorf("Expected %v, got %v", tc.expectedLen, len(subscriptions))
			}
		})
	}
}

func TestSubscriptionRepository_Update(t *testing.T) {
	type testCase struct {
		test                string
		initialSubscription entity.Subscription
		updatedSubscription entity.Subscription
		expectedErr         error
	}

	testCases := []testCase{
		{
			test:                "Update an existing subscription",
			initialSubscription: entity.Subscription{Name: "Test Subscription"},
			updatedSubscription: entity.Subscription{ID: 1, Name: "Updated Test Subscription"},
			expectedErr:         nil,
		},
		{
			test:                "Update a non-existing subscription",
			initialSubscription: entity.Subscription{Name: "Test Subscription"},
			updatedSubscription: entity.Subscription{ID: 3, Name: "Updated Test Subscription"},
			expectedErr:         repository.ErrUpdateSubscription,
		},
	}

	repo := memory.NewSubscriptionRepository()

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := repo.Create(context.Background(), tc.initialSubscription)
			if err != nil {
				t.Fatal(err)
			}

			updatedSubscription, err := repo.Update(context.Background(), tc.updatedSubscription)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("Expected error: %v, got: %v", tc.expectedErr, err)
			}

			if updatedSubscription != nil {
				found, err := repo.Get(context.Background(), updatedSubscription.ID)
				if err != nil {
					t.Fatal(err)
				}
				if updatedSubscription.Name != found.Name {
					t.Errorf("Expected %v, got %v", updatedSubscription.Name, found.Name)
				}
			}
		})
	}
}

func TestSubscriptionRepository_Delete(t *testing.T) {
	repo := memory.NewSubscriptionRepository()

	ctx := context.Background()

	t.Run("Non-existing subscription", func(t *testing.T) {
		err := repo.Delete(ctx, 3)
		if err == nil {
			t.Error("Expected error, got nil")
		}

		if !errors.Is(err, repository.ErrDeleteSubscription) {
			t.Errorf("Expected error: %v, got: %v", repository.ErrDeleteSubscription, err)
		}
	})

	t.Run("Existing subscription", func(t *testing.T) {
		_, err := repo.Create(context.Background(), entity.Subscription{Name: "Test Subscription"})
		if err != nil {
			t.Fatal(err)
		}

		err = repo.Delete(ctx, 1)
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}

		if _, err := repo.Get(context.Background(), 1); err == nil {
			t.Errorf("Expected subscription to be deleted")
		}
	})
}
