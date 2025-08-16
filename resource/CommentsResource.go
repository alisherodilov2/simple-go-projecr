package resource

import "github.com/alisherodilov2/go-first/models"

type CommentsResource struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func CommentCollection(comments []models.Comments) []CommentsResource {
	resources := make([]CommentsResource, len(comments))
	for i, comment := range comments {
		resources[i] = CommentsResource{
			ID:          comment.ID,
			Title:       comment.Title,
			Description: comment.Description,
		}
	}
	return resources
}
