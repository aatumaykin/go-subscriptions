package cycle_handler_test

import (
	"context"
	"testing"

	"git.home/alex/go-subscriptions/internal/api/handler/cycle_handler"
	"git.home/alex/go-subscriptions/internal/domain/repository"
	"git.home/alex/go-subscriptions/internal/domain/service"
	"git.home/alex/go-subscriptions/tests"
	"git.home/alex/go-subscriptions/tests/mock_repository"
	"git.home/alex/go-subscriptions/tests/tests_assert"
	"github.com/stretchr/testify/assert"
)

func TestGetCycles(t *testing.T) {
	type resp struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
		Days uint   `json:"days"`
	}

	testCases := []struct {
		name      string
		cycles    repository.Cycles
		mockError error
		expected  []resp
		wantErr   error
	}{
		{
			name: "Success",
			cycles: repository.Cycles{
				{ID: 1, Name: "Cycle 1", Days: 7},
				{ID: 2, Name: "Cycle 2", Days: 30},
			},
			mockError: nil,
			expected: []resp{
				{ID: 1, Name: "Cycle 1", Days: 7},
				{ID: 2, Name: "Cycle 2", Days: 30},
			},
		},
		{
			name:      "Error",
			cycles:    nil,
			mockError: tests.ErrTest,
			expected:  nil,
			wantErr:   tests.ErrTest,
		},
	}

	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(mock_repository.MockCycleRepository)
			mockRepo.On("GetAll", ctx).Return(tc.cycles, tc.mockError)

			cs := service.NewCycleService(mockRepo)

			response := cycle_handler.GetCycles(ctx, cs)(nil, nil)

			if tc.wantErr != nil {
				assert.ErrorIs(t, tc.wantErr, response.(error))
				return
			}

			tests_assert.EqualAsJSON(t, tc.expected, response)
		})
	}
}
