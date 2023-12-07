package cycle_handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"git.home/alex/go-subscriptions/internal/api/handler/cycle_handler"
	"git.home/alex/go-subscriptions/internal/domain/service"
	"git.home/alex/go-subscriptions/internal/repository/memory"
	"git.home/alex/go-subscriptions/tests/tests_assert"
	"github.com/stretchr/testify/assert"
)

func TestCreateCycle(t *testing.T) {
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
		name        string
		requestBody req
		expected    resp
		wantErr     error
	}{
		{
			name:        "Test Create Cycle",
			requestBody: req{Name: "Test Cycle", Days: 7},
			expected:    resp{ID: 1, Name: "Test Cycle", Days: 7},
		},
		{
			name:        "Test Create Cycle",
			requestBody: req{Name: "Test Cycle 2", Days: 30},
			expected:    resp{ID: 2, Name: "Test Cycle 2", Days: 30},
		},
		{
			name:        "Test validation error",
			requestBody: req{Name: "", Days: 0},
			expected:    resp{},
			wantErr:     service.ErrInvalidCycle,
		},
	}

	cs := service.NewCycleService(memory.NewCycleRepository())
	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			requestBodyBytes, _ := json.Marshal(tc.requestBody)
			r := &http.Request{
				Body: io.NopCloser(bytes.NewBuffer(requestBodyBytes)),
			}

			response := cycle_handler.CreateCycle(ctx, cs)(r, nil)

			if tc.wantErr != nil {
				assert.ErrorIs(t, tc.wantErr, response.(error))
				return
			}

			tests_assert.EqualAsJSON(t, tc.expected, response)
		})
	}
}
