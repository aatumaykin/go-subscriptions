package category_handler_test

import (
	"testing"

	"git.home/alex/go-subscriptions/internal/api/handler/category_handler"
	"git.home/alex/go-subscriptions/internal/domain/category/entity"
	"github.com/stretchr/testify/assert"
)

func TestCategoryToDTO(t *testing.T) {
	t.Run("Convert Category to DTO", func(t *testing.T) {
		category := entity.Category{
			ID:   1,
			Name: "Test Category",
		}
		expected := category_handler.CategoryDTO{
			ID:   1,
			Name: "Test Category",
		}

		result := category_handler.CategoryToDTO(category)

		assert.Equal(t, expected, result)
	})
}
