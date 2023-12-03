package service_test

import (
	"context"
	"testing"
	"time"

	. "git.home/alex/go-subscriptions/internal/domain/category/entity"
	. "git.home/alex/go-subscriptions/internal/domain/currency/entity"
	. "git.home/alex/go-subscriptions/internal/domain/cycle/entity"
	. "git.home/alex/go-subscriptions/internal/domain/subscription/entity"

	categoryservice "git.home/alex/go-subscriptions/internal/service/category"
	subscriptionservice "git.home/alex/go-subscriptions/internal/service/subscription"
)

var (
	categoryService     *categoryservice.Service
	subscriptionService *subscriptionservice.Service
)

func TestUseCase(t *testing.T) {
	ctx := context.Background()

	initServices(t)

	category := createCategory(ctx, t, "Media", 1)

	now := time.Now()
	nextMonth := now.Add(time.Hour * 24 * 30) // Add 30 days

	subscription := Subscription{
		ID:              1,
		Name:            "Yandex Plus",
		Category:        *category,
		Note:            "Yandex Plus subscription",
		Price:           3200,
		Currency:        RUB,
		Cycle:           Monthly,
		NextPaymentDate: PaymentDate(nextMonth),
	}

	createdSubscription, err := subscriptionService.Create(ctx, subscription)
	if err != nil {
		t.Fatal(err)
	}

	if *createdSubscription != subscription {
		t.Errorf("Expected subscription to be '%v', got %v", subscription, createdSubscription.Name)
	}
}

func initServices(t *testing.T) {
	t.Helper()

	var err error

	categoryService, err = categoryservice.NewCategoryService(categoryservice.WithMemoryCategoryRepository())
	if err != nil {
		t.Fatal(err)
	}

	subscriptionService, err = subscriptionservice.NewSubscriptionService(subscriptionservice.WithMemorySubscriptionRepository())
	if err != nil {
		t.Fatal(err)
	}
}

func createCategory(ctx context.Context, t *testing.T, name string, expectedID uint) *Category {
	t.Helper()

	category, err := categoryService.Create(ctx, name)
	if err != nil {
		t.Fatal(err)
	}
	if category.ID != expectedID {
		t.Errorf("Expected category ID to be %d, got %d", expectedID, category.ID)
	}
	if category.Name != name {
		t.Errorf("Expected category name to be '%s', got %s", name, category.Name)
	}

	return category
}
