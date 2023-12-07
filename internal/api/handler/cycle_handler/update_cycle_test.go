package cycle_handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
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

func TestUpdateCycle(t *testing.T) {
	type req struct {
		Name string `json:"name"`
		Days uint   `json:"days"`
	}

	type resp struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
		Days uint   `json:"days"`
	}

	testCases := []struct {
		name         string
		initialCycle entity.Cycle
		requestBody  req
		id           string
		expected     resp
		wantErr      error
	}{
		{
			name:         "success",
			initialCycle: entity.Cycle{Name: "Test Cycle", Days: 30},
			requestBody:  req{Name: "Updated name"},
			id:           "1",
			expected:     resp{ID: 1, Name: "Updated name", Days: 0},
		},
		{
			name:         "success",
			initialCycle: entity.Cycle{Name: "Test Cycle", Days: 30},
			requestBody:  req{Name: "Updated name", Days: 60},
			id:           "1",
			expected:     resp{ID: 1, Name: "Updated name", Days: 60},
		},
		{
			name:         "validation error",
			initialCycle: entity.Cycle{Name: "Test Cycle", Days: 30},
			requestBody:  req{Name: ""},
			id:           "1",
			wantErr:      service.ErrInvalidCycle,
		},
		{
			name:         "not found error",
			initialCycle: entity.Cycle{Name: "Test Cycle", Days: 30},
			requestBody:  req{Name: "Updated name"},
			id:           "10",
			wantErr:      repository.ErrNotFoundCycle,
		},
	}

	cs := service.NewCycleService(memory.NewCycleRepository())
	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, _ = cs.CreateCycle(ctx, tc.initialCycle)

			requestBodyBytes, _ := json.Marshal(tc.requestBody)
			r := &http.Request{
				Body: io.NopCloser(bytes.NewBuffer(requestBodyBytes)),
			}
			ps := httprouter.Params{{Key: "id", Value: tc.id}}

			response := cycle_handler.UpdateCycle(ctx, cs)(r, ps)

			if tc.wantErr != nil {
				assert.ErrorIs(t, tc.wantErr, response.(error))
				return
			}

			tests_assert.EqualAsJSON(t, tc.expected, response)
		})
	}
}
