package resource

import "github.com/alisherodilov2/go-first/models"

type UserResource struct {
	ID       uint   `json:"id"`
	Username string `json:"name"`
	Password string `json:"password"`
}

func UserMake(user models.User) UserResource {
	return UserResource{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
	}
}
