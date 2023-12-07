package category_handler

import (
	"context"
	"encoding/json"
	"net/http"

	"git.home/alex/go-subscriptions/internal/api/api_response"
	"git.home/alex/go-subscriptions/internal/api/handler/error_handler"
	"git.home/alex/go-subscriptions/internal/api/middleware"
	"git.home/alex/go-subscriptions/internal/domain/entity"
	"git.home/alex/go-subscriptions/internal/domain/service"
	"github.com/julienschmidt/httprouter"
)

func CreateHandle(ctx context.Context, categoryService *service.CategoryService) httprouter.Handle {
	return middleware.SetJSONContentType(func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var requestDTO struct {
			Name string `json:"name"`
		}
		err := json.NewDecoder(r.Body).Decode(&requestDTO)
		if err != nil {
			error_handler.HandleError(w, err)
			return
		}

		createdCategory, err := categoryService.CreateCategory(ctx, entity.Category{
			Name: requestDTO.Name,
		})
		if err != nil {
			error_handler.HandleError(w, err)
			return
		}

		type responseDTO struct {
			ID   uint   `json:"id"`
			Name string `json:"name"`
		}

		response, err := api_response.Success(responseDTO{
			ID:   createdCategory.ID,
			Name: createdCategory.Name,
		}).ToJSON()
		if err != nil {
			error_handler.HandleError(w, err)
			return
		}

		_, _ = w.Write(response)
	})
}
