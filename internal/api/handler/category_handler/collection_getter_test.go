package category_handler_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"git.home/alex/go-subscriptions/internal/api/api_response"
	"git.home/alex/go-subscriptions/internal/api/handler/category_handler"
	"git.home/alex/go-subscriptions/internal/domain/repository"
	"git.home/alex/go-subscriptions/internal/domain/service"
	"git.home/alex/go-subscriptions/tests"
	"git.home/alex/go-subscriptions/tests/mock_repository"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestCollectionGetterHandle(t *testing.T) {
	testCases := []struct {
		name           string
		categories     repository.Categories
		mockError      error
		expectedStatus int
		expectedBody   api_response.ResponseDTO
	}{
		{
			name: "Success",
			categories: repository.Categories{
				{ID: 1, Name: "Category 1"},
				{ID: 2, Name: "Category 2"},
			},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody: api_response.ResponseDTO{
				Status: "success",
				Data: []category_handler.CategoryDTO{
					{ID: 1, Name: "Category 1"},
					{ID: 2, Name: "Category 2"},
				},
			},
		},
		{
			name:           "Error",
			categories:     nil,
			mockError:      tests.ErrTest,
			expectedStatus: http.StatusOK,
			expectedBody: api_response.ResponseDTO{
				Status: "error",
				Error:  "some error",
			},
		},
	}

	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(mock_repository.MockCategoryRepository)
			mockRepo.On("GetAll", context.Background()).Return(tc.categories, tc.mockError)

			cs := service.NewCategoryService(mockRepo)

			router := httprouter.New()
			router.GET("/categories", category_handler.CollectionGetterHandle(ctx, cs))

			req, _ := http.NewRequestWithContext(ctx, "GET", "/categories", nil)
			rr := httptest.NewRecorder()

			router.ServeHTTP(rr, req)

			assert.Equal(t, tc.expectedStatus, rr.Code)
			assert.Equal(t, "application/json", rr.Header().Get("Content-Type"), "handler returned wrong content type")

			expectedBody, err := json.Marshal(tc.expectedBody)
			assert.NoError(t, err)

			assert.Equal(t, string(expectedBody), rr.Body.String())
		})
	}
}
