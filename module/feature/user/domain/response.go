package domain

import (
	"fungicheck/module/entities"
	"time"
)

type UserResponse struct {
	ID        uint64    `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Name      string    `json:"name"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

func UserFormatter(user *entities.UserModels) *UserResponse {
	result := &UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Password:  "",
		Name:      user.Name,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
	}
	return result
}
