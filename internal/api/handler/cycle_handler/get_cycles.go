package cycle_handler

import (
	"context"
	"net/http"

	"git.home/alex/go-subscriptions/internal/api/api_response"
	"git.home/alex/go-subscriptions/internal/domain/service"
	"github.com/julienschmidt/httprouter"
)

func GetCycles(ctx context.Context, cs *service.CycleService) api_response.Handle {
	return func(_ *http.Request, _ httprouter.Params) any {
		cycles, err := cs.GetAllCycles(ctx)
		if err != nil {
			return err
		}

		type resp struct {
			ID   uint   `json:"id"`
			Name string `json:"name"`
			Days uint   `json:"days"`
		}

		cyclesResp := make([]resp, len(cycles))
		for i, cycle := range cycles {
			cyclesResp[i] = resp{
				ID:   cycle.ID,
				Name: cycle.Name,
				Days: cycle.Days,
			}
		}

		return cyclesResp
	}
}
