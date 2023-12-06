package health_handler

import (
	"encoding/json"
	"net/http"

	"git.home/alex/go-subscriptions/internal/version"
	"github.com/julienschmidt/httprouter"
)

type ResponseDTO struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}

func Handle() httprouter.Handle {
	return func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		dto := ResponseDTO{
			Status:  "pass",
			Version: version.Version,
		}

		response, err := json.Marshal(dto)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(response)
	}
}
