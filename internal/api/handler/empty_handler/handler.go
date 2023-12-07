package empty_handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"git.home/alex/go-subscriptions/internal/api/api_response"
	"github.com/julienschmidt/httprouter"
)

var (
	ErrNotImplemented = errors.New("not implemented")
)

func Handle() httprouter.Handle {
	return func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		dto := api_response.Error(ErrNotImplemented)
		response, _ := json.Marshal(dto)

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(response)
	}
}
