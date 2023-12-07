package category_handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
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

func TestUpdateCategoryHandle(t *testing.T) {
	type requestDTO struct {
		Name string `json:"name"`
	}

	type responseDTO struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	testCases := []struct {
		name            string
		initialCategory entity.Category
		requestBody     requestDTO
		id              string
		expectedStatus  int
		expectedBody    api_response.ResponseDTO
	}{
		{
			name:            "success",
			initialCategory: entity.Category{Name: "Test Category"},
			requestBody:     requestDTO{Name: "Updated Category"},
			id:              "1",
			expectedStatus:  http.StatusOK,
			expectedBody:    api_response.Success(responseDTO{ID: 1, Name: "Updated Category"}),
		},
		{
			name:            "success",
			initialCategory: entity.Category{Name: "Test Category 2"},
			requestBody:     requestDTO{Name: "Updated Category 2"},
			id:              "2",
			expectedStatus:  http.StatusOK,
			expectedBody:    api_response.Success(responseDTO{ID: 2, Name: "Updated Category 2"}),
		},
		{
			name:            "validation error",
			initialCategory: entity.Category{Name: "Test Category 3"},
			requestBody:     requestDTO{Name: ""},
			id:              "3",
			expectedStatus:  http.StatusOK,
			expectedBody:    api_response.Error(service.ErrInvalidCategory),
		},
		{
			name:            "not found error",
			initialCategory: entity.Category{Name: "Test Category 4"},
			requestBody:     requestDTO{Name: "Updated Category 2"},
			id:              "10",
			expectedStatus:  http.StatusOK,
			expectedBody:    api_response.Error(repository.ErrNotFoundCategory),
		},
	}

	cs := service.NewCategoryService(memory.NewCategoryRepository())
	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := cs.CreateCategory(ctx, tc.initialCategory)
			assert.NoError(t, err)

			requestBodyBytes, _ := json.Marshal(tc.requestBody)

			w := httptest.NewRecorder()
			r := &http.Request{
				Body: io.NopCloser(bytes.NewBuffer(requestBodyBytes)),
			}
			ps := httprouter.Params{{Key: "id", Value: tc.id}}

			category_handler.UpdateHandle(ctx, cs)(w, r, ps)

			assert.Equal(t, tc.expectedStatus, w.Code)
			assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

			expectedBody, err := json.Marshal(tc.expectedBody)
			assert.NoError(t, err)

			assert.Equal(t, string(expectedBody), w.Body.String())
		})
	}
}
