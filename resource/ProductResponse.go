package resource

import (
	"github.com/alisherodilov2/go-first/models"
)

type ProductResponse struct {
	ID          uint               `json:"id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Image       string             `json:"image"`
	Comments    []CommentsResource `json:"comments"`
}

func ProductCollection(products []models.Products) []ProductResponse {
	resources := make([]ProductResponse, len(products))
	for i, product := range products {
		resources[i] = ProductResponse{
			ID:          product.ID,
			Title:       product.Title,
			Description: product.Description,
			Image:       product.Image,
			Comments:    CommentCollection(product.Comments),
		}
	}
	return resources
}

func ProductMake(product models.Products) ProductResponse {
	return ProductResponse{
		ID:          product.ID,
		Title:       product.Title,
		Description: product.Description,
		Image:       product.Image,
		Comments:    CommentCollection(product.Comments),
	}
}
