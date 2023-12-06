package category_handler

import "git.home/alex/go-subscriptions/internal/domain/entity"

type CategoryDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func CategoryToDTO(category entity.Category) CategoryDTO {
	return CategoryDTO{
		ID:   category.ID,
		Name: category.Name,
	}
}
