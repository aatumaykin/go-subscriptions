package category_handler

import (
	"context"
	"net/http"

	"git.home/alex/go-subscriptions/internal/api/api_response"
	"git.home/alex/go-subscriptions/internal/api/handler/error_handler"
	"git.home/alex/go-subscriptions/internal/api/middleware"
	"git.home/alex/go-subscriptions/internal/domain/service"
	"github.com/julienschmidt/httprouter"
)

func CollectionGetHandle(ctx context.Context, categoryService *service.CategoryService) httprouter.Handle {
	return middleware.SetJSONContentType(func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		categories, err := categoryService.GetAllCategories(ctx)
		if err != nil {
			error_handler.HandleError(w, err)

			return
		}

		type responseDTO struct {
			ID   uint   `json:"id"`
			Name string `json:"name"`
		}

		categoryDTOs := make([]responseDTO, len(categories))
		for i, category := range categories {
			categoryDTOs[i] = responseDTO{
				ID:   category.ID,
				Name: category.Name,
			}
		}

		response, err := api_response.Success(categoryDTOs).ToJSON()
		if err != nil {
			error_handler.HandleError(w, err)

			return
		}

		_, _ = w.Write(response)
	})
}
