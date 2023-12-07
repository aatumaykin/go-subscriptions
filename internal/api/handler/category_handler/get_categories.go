package category_handler

import (
	"context"
	"net/http"

	"git.home/alex/go-subscriptions/internal/api/api_response"
	"git.home/alex/go-subscriptions/internal/domain/service"
	"github.com/julienschmidt/httprouter"
)

func GetCategories(ctx context.Context, cs *service.CategoryService) api_response.Handle {
	return func(_ *http.Request, _ httprouter.Params) any {
		categories, err := cs.GetAllCategories(ctx)
		if err != nil {
			return err
		}

		type resp struct {
			ID   uint   `json:"id"`
			Name string `json:"name"`
		}

		categoryDTOs := make([]resp, len(categories))
		for i, category := range categories {
			categoryDTOs[i] = resp{
				ID:   category.ID,
				Name: category.Name,
			}
		}

		return categoryDTOs
	}
}
