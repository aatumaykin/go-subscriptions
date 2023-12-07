package category_handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"git.home/alex/go-subscriptions/internal/api/api_response"
	"git.home/alex/go-subscriptions/internal/domain/service"
	"github.com/julienschmidt/httprouter"
)

func UpdateCategory(ctx context.Context, cs *service.CategoryService) api_response.Handle {
	return func(r *http.Request, ps httprouter.Params) any {
		id, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			return err
		}

		var req struct {
			Name string `json:"name"`
		}
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			return err
		}

		category, err := cs.GetCategory(ctx, uint(id))
		if err != nil {
			return err
		}

		category.Name = req.Name
		updatedCategory, err := cs.UpdateCategory(ctx, *category)
		if err != nil {
			return err
		}

		type resp struct {
			ID   uint   `json:"id"`
			Name string `json:"name"`
		}

		return resp{
			ID:   updatedCategory.ID,
			Name: updatedCategory.Name,
		}
	}
}
