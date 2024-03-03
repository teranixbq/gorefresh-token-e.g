package response

import "refresh/internal/user/model"

type ResponseLogin struct {
	Username     string `json:"username,omitempty"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type ResponseUser struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

func ModelToResponseLogin(username, accessToken, refreshToken string) ResponseLogin {
	return ResponseLogin{
		Username:     username,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}

func ModelToResponseUser(data model.Users) ResponseUser {
	return ResponseUser{
		Id:       data.Id,
		Username: data.Username,
	}
}