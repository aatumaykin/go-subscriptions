package category_handler

import "git.home/alex/go-subscriptions/internal/domain/category/entity"

type categoryDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func categoryToDTO(category entity.Category) categoryDTO {
	return categoryDTO{
		ID:   category.ID,
		Name: category.Name,
	}
}
