package empty_handler

import (
	"errors"
	"net/http"

	"git.home/alex/go-subscriptions/internal/api/api_response"
	"git.home/alex/go-subscriptions/internal/api/handler/error_handler"
	"git.home/alex/go-subscriptions/internal/api/middleware"
	"github.com/julienschmidt/httprouter"
)

var (
	ErrNotImplemented = errors.New("not implemented")
)

func Handle() httprouter.Handle {
	return middleware.SetJSONContentType(func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {

		response, err := api_response.Error(ErrNotImplemented).ToJSON()
		if err != nil {
			error_handler.HandleError(w, err)
			return
		}
		_, _ = w.Write(response)
	})
}
