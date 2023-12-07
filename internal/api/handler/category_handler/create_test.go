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
	"git.home/alex/go-subscriptions/internal/domain/service"
	"git.home/alex/go-subscriptions/internal/repository/memory"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestCreateCategoryHandle(t *testing.T) {
	type requestDTO struct {
		Name string `json:"name"`
	}

	type responseDTO struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	testCases := []struct {
		name           string
		requestBody    requestDTO
		expectedStatus int
		expectedBody   api_response.ResponseDTO
	}{
		{
			name:           "success",
			requestBody:    requestDTO{Name: "Test Category"},
			expectedStatus: http.StatusOK,
			expectedBody:   api_response.Success(responseDTO{ID: 1, Name: "Test Category"}),
		},
		{
			name:           "success",
			requestBody:    requestDTO{Name: "Test Category 2"},
			expectedStatus: http.StatusOK,
			expectedBody:   api_response.Success(responseDTO{ID: 2, Name: "Test Category 2"}),
		},
		{
			name:           "error",
			requestBody:    requestDTO{Name: ""},
			expectedStatus: http.StatusOK,
			expectedBody:   api_response.Error(service.ErrInvalidCategory),
		},
	}

	cs := service.NewCategoryService(memory.NewCategoryRepository())
	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			requestBodyBytes, _ := json.Marshal(tc.requestBody)

			w := httptest.NewRecorder()
			r := &http.Request{
				Body: io.NopCloser(bytes.NewBuffer(requestBodyBytes)),
			}
			ps := httprouter.Params{}

			category_handler.CreateHandle(ctx, cs)(w, r, ps)

			assert.Equal(t, tc.expectedStatus, w.Code)
			assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

			expectedBody, err := json.Marshal(tc.expectedBody)
			assert.NoError(t, err)

			assert.Equal(t, string(expectedBody), w.Body.String())
		})
	}
}
