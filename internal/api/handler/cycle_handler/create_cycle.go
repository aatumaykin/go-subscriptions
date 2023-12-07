package cycle_handler

import (
	"context"
	"encoding/json"
	"net/http"

	"git.home/alex/go-subscriptions/internal/api/api_response"
	"git.home/alex/go-subscriptions/internal/domain/entity"
	"git.home/alex/go-subscriptions/internal/domain/service"
	"github.com/julienschmidt/httprouter"
)

func CreateCycle(ctx context.Context, cs *service.CycleService) api_response.Handle {
	return func(r *http.Request, _ httprouter.Params) any {
		var req struct {
			Name string `json:"name"`
			Days uint   `json:"days"`
		}

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			return err
		}

		createdCycle, err := cs.CreateCycle(ctx, entity.Cycle{
			Name: req.Name,
			Days: req.Days,
		})
		if err != nil {
			return err
		}

		type resp struct {
			ID   uint   `json:"id"`
			Name string `json:"name"`
			Days uint   `json:"days"`
		}

		return resp{
			ID:   createdCycle.ID,
			Name: createdCycle.Name,
			Days: createdCycle.Days,
		}
	}
}
