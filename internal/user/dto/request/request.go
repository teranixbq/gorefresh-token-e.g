package request

import "refresh/internal/user/model"

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RequestRefresh struct {
	RefreshToken string `json:"refresh_token"`
}

func UserRequestToModel(data UserRequest) model.Users {
	return model.Users{
		Username: data.Username,
		Password: data.Password,
	}
}