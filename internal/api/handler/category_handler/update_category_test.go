package category_handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"git.home/alex/go-subscriptions/internal/api/handler/category_handler"
	"git.home/alex/go-subscriptions/internal/domain/entity"
	"git.home/alex/go-subscriptions/internal/domain/repository"
	"git.home/alex/go-subscriptions/internal/domain/service"
	"git.home/alex/go-subscriptions/internal/repository/memory"
	"git.home/alex/go-subscriptions/tests/tests_assert"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestUpdateCategory(t *testing.T) {
	type req struct {
		Name string `json:"name"`
	}

	type resp struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	testCases := []struct {
		name            string
		initialCategory entity.Category
		requestBody     req
		id              string
		expected        resp
		wantErr         error
	}{
		{
			name:            "success",
			initialCategory: entity.Category{Name: "Test Category"},
			requestBody:     req{Name: "Updated Category"},
			id:              "1",
			expected:        resp{ID: 1, Name: "Updated Category"},
		},
		{
			name:            "success",
			initialCategory: entity.Category{Name: "Test Category 2"},
			requestBody:     req{Name: "Updated Category 2"},
			id:              "2",
			expected:        resp{ID: 2, Name: "Updated Category 2"},
		},
		{
			name:            "validation error",
			initialCategory: entity.Category{Name: "Test Category 3"},
			requestBody:     req{Name: ""},
			id:              "3",
			wantErr:         service.ErrInvalidCategory,
		},
		{
			name:            "not found error",
			initialCategory: entity.Category{Name: "Test Category 4"},
			requestBody:     req{Name: "Updated Category 2"},
			id:              "10",
			wantErr:         repository.ErrNotFoundCategory,
		},
	}

	cs := service.NewCategoryService(memory.NewCategoryRepository())
	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, _ = cs.CreateCategory(ctx, tc.initialCategory)

			requestBodyBytes, _ := json.Marshal(tc.requestBody)
			r := &http.Request{
				Body: io.NopCloser(bytes.NewBuffer(requestBodyBytes)),
			}
			ps := httprouter.Params{{Key: "id", Value: tc.id}}

			response := category_handler.UpdateCategory(ctx, cs)(r, ps)

			if tc.wantErr != nil {
				assert.ErrorIs(t, tc.wantErr, response.(error))
				return
			}

			tests_assert.EqualAsJSON(t, tc.expected, response)
		})
	}
}
