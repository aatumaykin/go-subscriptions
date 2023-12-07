package category_handler

import (
	"context"
	"net/http"
	"strconv"

	"git.home/alex/go-subscriptions/internal/api/api_response"
	"git.home/alex/go-subscriptions/internal/api/handler/error_handler"
	"git.home/alex/go-subscriptions/internal/api/middleware"
	"git.home/alex/go-subscriptions/internal/domain/service"
	"github.com/julienschmidt/httprouter"
)

func DeleteHandle(ctx context.Context, categoryService *service.CategoryService) httprouter.Handle {
	return middleware.SetJSONContentType(func(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
		id, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			error_handler.HandleError(w, err)
			return
		}

		err = categoryService.DeleteCategory(ctx, uint(id))
		if err != nil {
			error_handler.HandleError(w, err)
			return
		}

		response, err := api_response.Success("OK").ToJSON()
		if err != nil {
			error_handler.HandleError(w, err)
			return
		}

		_, _ = w.Write(response)
	})
}
