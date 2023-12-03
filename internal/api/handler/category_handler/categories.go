package category_handler

import (
	"context"
	"encoding/json"
	"net/http"

	categoryserviceinterface "git.home/alex/go-subscriptions/internal/domain/category/service"
	"github.com/julienschmidt/httprouter"
)

func AllCategoriesHandle(ctx context.Context, categoryService categoryserviceinterface.CategoryService) httprouter.Handle {
	return func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		categories, err := categoryService.GetAll(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Convert the categories to a slice of CategoryDTOs
		categoryDTOs := make([]categoryDTO, len(categories))
		for i, category := range categories {
			categoryDTOs[i] = categoryToDTO(category)
		}

		response, err := json.Marshal(categoryDTOs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(response)
	}
}
