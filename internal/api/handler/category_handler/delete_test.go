package category_handler_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"git.home/alex/go-subscriptions/internal/api/api_response"
	"git.home/alex/go-subscriptions/internal/api/handler/category_handler"
	"git.home/alex/go-subscriptions/internal/domain/entity"
	"git.home/alex/go-subscriptions/internal/domain/repository"
	"git.home/alex/go-subscriptions/internal/domain/service"
	"git.home/alex/go-subscriptions/internal/repository/memory"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestDeleteHandle(t *testing.T) {
	testCases := []struct {
		name           string
		id             string
		category       entity.Category
		expectedStatus int
		expectedBody   api_response.ResponseDTO
	}{
		{
			name:           "success",
			id:             "1",
			category:       entity.Category{ID: 1, Name: "Test Category"},
			expectedStatus: http.StatusOK,
			expectedBody:   api_response.Success("OK"),
		},
		{
			name:           "error",
			id:             "10",
			category:       entity.Category{ID: 2, Name: "Test Category"},
			expectedStatus: http.StatusOK,
			expectedBody:   api_response.Error(repository.ErrNotFoundCategory),
		},
	}

	categoryService := service.NewCategoryService(memory.NewCategoryRepository())
	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := categoryService.CreateCategory(ctx, tc.category)
			assert.NoError(t, err)

			w := httptest.NewRecorder()
			r := &http.Request{}
			ps := httprouter.Params{{Key: "id", Value: tc.id}}

			category_handler.DeleteHandle(ctx, categoryService)(w, r, ps)

			assert.Equal(t, tc.expectedStatus, w.Code)
			assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

			expectedBody, err := json.Marshal(tc.expectedBody)
			assert.NoError(t, err)

			assert.Equal(t, string(expectedBody), w.Body.String())
		})
	}
}
