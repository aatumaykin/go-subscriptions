package category_handler

import (
	"context"
	"net/http"

	"git.home/alex/go-subscriptions/internal/api/api_response"
	"git.home/alex/go-subscriptions/internal/api/handler/error_handler"
	"git.home/alex/go-subscriptions/internal/domain/service"
	"github.com/julienschmidt/httprouter"
)

func CollectionGetterHandle(ctx context.Context, categoryService service.CategoryService) httprouter.Handle {
	return func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		categories, err := categoryService.GetAllCategories(ctx)
		if err != nil {
			error_handler.HandleError(w, err)

			return
		}

		// Convert the categories to a slice of CategoryDTOs
		categoryDTOs := make([]CategoryDTO, len(categories))
		for i, category := range categories {
			categoryDTOs[i] = CategoryToDTO(category)
		}

		responseDto := api_response.SuccessResponse
		responseDto.Data = categoryDTOs
		response, err := responseDto.ToJSON()
		if err != nil {
			error_handler.HandleError(w, err)

			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(response)
	}
}
