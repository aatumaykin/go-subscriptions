package cycle_handler_test

import (
	"context"
	"testing"

	"git.home/alex/go-subscriptions/internal/api/handler/cycle_handler"
	"git.home/alex/go-subscriptions/internal/domain/entity"
	"git.home/alex/go-subscriptions/internal/domain/repository"
	"git.home/alex/go-subscriptions/internal/domain/service"
	"git.home/alex/go-subscriptions/internal/repository/memory"
	"git.home/alex/go-subscriptions/tests/tests_assert"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestGetCycle(t *testing.T) {
	type resp struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
		Days uint   `json:"days"`
	}

	testCases := []struct {
		name     string
		id       string
		cycle    entity.Cycle
		expected resp
		wantErr  error
	}{
		{
			name:     "success",
			id:       "1",
			cycle:    entity.Cycle{ID: 1, Name: "Test Cycle", Days: 7},
			expected: resp{ID: 1, Name: "Test Cycle", Days: 7},
		},
		{
			name:    "error",
			id:      "10",
			cycle:   entity.Cycle{ID: 2, Name: "Test Cycle", Days: 7},
			wantErr: repository.ErrNotFoundCycle,
		},
	}

	cs := service.NewCycleService(memory.NewCycleRepository())
	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, _ = cs.CreateCycle(ctx, tc.cycle)
			ps := httprouter.Params{{Key: "id", Value: tc.id}}

			response := cycle_handler.GetCycle(ctx, cs)(nil, ps)

			if tc.wantErr != nil {
				assert.ErrorIs(t, tc.wantErr, response.(error))
				return
			}

			tests_assert.EqualAsJSON(t, tc.expected, response)
		})
	}
}
