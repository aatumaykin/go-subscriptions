package category_handler

import (
	"context"
	"encoding/json"
	"net/http"

	"git.home/alex/go-subscriptions/internal/api/api_response"
	"git.home/alex/go-subscriptions/internal/domain/entity"
	"git.home/alex/go-subscriptions/internal/domain/service"
	"github.com/julienschmidt/httprouter"
)

func CreateCategory(ctx context.Context, cs *service.CategoryService) api_response.Handle {
	return func(r *http.Request, _ httprouter.Params) any {
		var req struct {
			Name string `json:"name"`
		}

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			return err
		}

		createdCategory, err := cs.CreateCategory(ctx, entity.Category{
			Name: req.Name,
		})
		if err != nil {
			return err
		}

		type resp struct {
			ID   uint   `json:"id"`
			Name string `json:"name"`
		}

		return resp{
			ID:   createdCategory.ID,
			Name: createdCategory.Name,
		}
	}
}
