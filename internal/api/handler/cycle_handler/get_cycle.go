package cycle_handler

import (
	"context"
	"net/http"
	"strconv"

	"git.home/alex/go-subscriptions/internal/api/api_response"
	"git.home/alex/go-subscriptions/internal/domain/service"
	"github.com/julienschmidt/httprouter"
)

func GetCycle(ctx context.Context, cs *service.CycleService) api_response.Handle {
	return func(_ *http.Request, ps httprouter.Params) any {
		id, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			return err
		}

		cycle, err := cs.GetCycle(ctx, uint(id))
		if err != nil {
			return err
		}

		type resp struct {
			ID   uint   `json:"id"`
			Name string `json:"name"`
			Days uint   `json:"days"`
		}

		return resp{
			ID:   cycle.ID,
			Name: cycle.Name,
			Days: cycle.Days,
		}
	}
}
