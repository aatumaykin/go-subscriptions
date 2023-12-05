package category_handler_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"git.home/alex/go-subscriptions/internal/api/api_response"
	"git.home/alex/go-subscriptions/internal/api/handler/category_handler"
	"git.home/alex/go-subscriptions/internal/domain/category/repository"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

type MockCollectionGetter struct {
}

// GetAll is a mock method for the GetAll function.
func (m *MockCollectionGetter) GetAll(ctx context.Context) (repository.Categories, error) {
	return repository.Categories{
		{ID: 1, Name: "Category 1"},
		{ID: 2, Name: "Category 2"},
	}, nil
}

func TestCollectionGetterHandle(t *testing.T) {
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		categoryService := new(MockCollectionGetter)

		handle := category_handler.CollectionGetterHandle(ctx, categoryService)

		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/categories", nil)
		p := httprouter.Params{}

		handle(w, r, p)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

		responseDto := api_response.SuccessResponse
		responseDto.Data = []category_handler.CategoryDTO{
			{ID: 1, Name: "Category 1"},
			{ID: 2, Name: "Category 2"},
		}
		expectedResponse, _ := responseDto.ToJSON()

		assert.Equal(t, string(expectedResponse), w.Body.String())
	})
}
