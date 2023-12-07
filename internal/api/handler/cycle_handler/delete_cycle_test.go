package cycle_handler_test

import (
	"context"
	"testing"

	"git.home/alex/go-subscriptions/internal/api/handler/cycle_handler"
	"git.home/alex/go-subscriptions/internal/domain/entity"
	"git.home/alex/go-subscriptions/internal/domain/repository"
	"git.home/alex/go-subscriptions/internal/domain/service"
	"git.home/alex/go-subscriptions/internal/repository/memory"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestDeleteCycle(t *testing.T) {
	testCases := []struct {
		name    string
		id      string
		cycle   entity.Cycle
		wantErr error
	}{
		{
			name:    "success",
			id:      "1",
			cycle:   entity.Cycle{ID: 1, Name: "Test Cycle", Days: 30},
			wantErr: nil,
		},
		{
			name:    "error",
			id:      "10",
			cycle:   entity.Cycle{ID: 2, Name: "Test Cycle", Days: 30},
			wantErr: repository.ErrNotFoundCycle,
		},
	}

	cs := service.NewCycleService(memory.NewCycleRepository())
	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, _ = cs.CreateCycle(ctx, tc.cycle)
			ps := httprouter.Params{{Key: "id", Value: tc.id}}

			response := cycle_handler.DeleteCycle(ctx, cs)(nil, ps)

			if tc.wantErr != nil {
				assert.ErrorIs(t, tc.wantErr, response.(error))
				return
			}

			assert.Nil(t, response)
		})
	}
}
