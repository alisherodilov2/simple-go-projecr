package resource

import "github.com/alisherodilov2/go-first/models"

type BookResource struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func BookCollection(books []models.Book) []BookResource {
	resources := make([]BookResource, len(books))
	for i, book := range books {
		resources[i] = BookResource{
			ID:          book.ID,
			Title:       book.Title,
			Author:      book.Author,
			Description: book.Description,
		}
	}
	return resources
}
