package domain

import (
	"fastfooducate/module/entities"
	"time"
)

type UserResponse struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

type UserLoginResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type LoginResponse struct {
	User        UserLoginResponse `json:"user"`
	AccessToken string            `json:"access_token"`
}

func LoginFormatter(user *entities.UserModels, accessToken string) LoginResponse {
	return LoginResponse{
		User: UserLoginResponse{
			Name:  user.Name,
			Email: user.Email,
		},
		AccessToken: accessToken,
	}
}

func RegisterFormatter(user *entities.UserModels) *UserResponse {
	userFormatter := &UserResponse{
		Email: user.Email,
		Role:  user.Role,
		Name:  user.Name,
	}
	return userFormatter
}
