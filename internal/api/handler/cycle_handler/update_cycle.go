package cycle_handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"git.home/alex/go-subscriptions/internal/api/api_response"
	"git.home/alex/go-subscriptions/internal/domain/service"
	"github.com/julienschmidt/httprouter"
)

func UpdateCycle(ctx context.Context, cs *service.CycleService) api_response.Handle {
	return func(r *http.Request, ps httprouter.Params) any {
		id, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			return err
		}

		var req struct {
			Name string `json:"name"`
			Days uint   `json:"days"`
		}
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			return err
		}

		cycle, err := cs.GetCycle(ctx, uint(id))
		if err != nil {
			return err
		}

		cycle.Name = req.Name
		cycle.Days = req.Days
		updatedCycle, err := cs.UpdateCycle(ctx, *cycle)
		if err != nil {
			return err
		}

		type resp struct {
			ID   uint   `json:"id"`
			Name string `json:"name"`
			Days uint   `json:"days"`
		}

		return resp{
			ID:   updatedCycle.ID,
			Name: updatedCycle.Name,
			Days: updatedCycle.Days,
		}
	}
}
