package category_handler

import (
	"context"
	"net/http"
	"strconv"

	"git.home/alex/go-subscriptions/internal/api/api_response"
	"git.home/alex/go-subscriptions/internal/api/handler/error_handler"
	"git.home/alex/go-subscriptions/internal/api/middleware"
	"git.home/alex/go-subscriptions/internal/domain/service"
	"github.com/julienschmidt/httprouter"
)

func GetHandle(ctx context.Context, categoryService *service.CategoryService) httprouter.Handle {
	return middleware.SetJSONContentType(func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		id, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			error_handler.HandleError(w, err)
			return
		}

		category, err := categoryService.GetCategory(ctx, uint(id))
		if err != nil {
			error_handler.HandleError(w, err)
			return
		}

		type responseDTO struct {
			ID   uint   `json:"id"`
			Name string `json:"name"`
		}

		response, err := api_response.Success(responseDTO{
			ID:   category.ID,
			Name: category.Name,
		}).ToJSON()
		if err != nil {
			error_handler.HandleError(w, err)
			return
		}

		_, _ = w.Write(response)
	})
}
