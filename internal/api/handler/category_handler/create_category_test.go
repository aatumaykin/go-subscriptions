package category_handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"git.home/alex/go-subscriptions/internal/api/handler/category_handler"
	"git.home/alex/go-subscriptions/internal/domain/service"
	"git.home/alex/go-subscriptions/internal/repository/memory"
	"git.home/alex/go-subscriptions/tests/tests_assert"
	"github.com/stretchr/testify/assert"
)

func TestCreateCategory(t *testing.T) {
	type req struct {
		Name string `json:"name"`
	}

	type resp struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	testCases := []struct {
		name        string
		requestBody req
		expected    resp
		wantErr     error
	}{
		{
			name:        "Test Create Category",
			requestBody: req{Name: "Test Category"},
			expected:    resp{ID: 1, Name: "Test Category"},
		},
		{
			name:        "Test Create Category",
			requestBody: req{Name: "Test Category 2"},
			expected:    resp{ID: 2, Name: "Test Category 2"},
		},
		{
			name:        "Test validation error",
			requestBody: req{Name: ""},
			expected:    resp{},
			wantErr:     service.ErrInvalidCategory,
		},
	}

	cs := service.NewCategoryService(memory.NewCategoryRepository())
	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			requestBodyBytes, _ := json.Marshal(tc.requestBody)
			r := &http.Request{
				Body: io.NopCloser(bytes.NewBuffer(requestBodyBytes)),
			}

			response := category_handler.CreateCategory(ctx, cs)(r, nil)

			if tc.wantErr != nil {
				assert.ErrorIs(t, tc.wantErr, response.(error))
				return
			}

			tests_assert.EqualAsJSON(t, tc.expected, response)
		})
	}
}
