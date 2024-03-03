package service

import (
	"errors"

	"refresh/internal/app/config"
	"refresh/internal/user/dto/request"
	"refresh/internal/user/dto/response"
	"refresh/pkg/auth"
	"refresh/pkg/hash"

	"refresh/internal/user/repository"
)

type userService struct {
	userRepository repository.UserRepositoryInterface
	auth           auth.Token
}

type UserServiceInterface interface {
	InsertUser(data request.UserRequest) error
	Login(data request.UserRequest) (response.ResponseLogin, error)
	RefreshToken(id string) (response.ResponseLogin, error)
	SelectUserById(username string) (response.ResponseUser, error)
}

func NewUserService(userRepository repository.UserRepositoryInterface, auth auth.Token) UserServiceInterface {
	return &userService{
		userRepository: userRepository,
		auth: auth,
	}
}

func (user *userService) InsertUser(data request.UserRequest) error {

	password, err := hash.HashPass(data.Password)
	if err != nil {
		return err
	}

	data.Password = password
	err = user.userRepository.InsertUser(data)
	if err != nil {
		return err
	}

	return nil
}

func (user *userService) Login(data request.UserRequest) (response.ResponseLogin, error) {

	dataUser, err := user.userRepository.SelectUsername(data.Username)
	if err != nil {
		return response.ResponseLogin{}, err
	}

	if !hash.CompareHash(dataUser.Password, data.Password) {
		return response.ResponseLogin{}, errors.New("wrong password")
	}

	accesToken, err := user.auth.CreateAccessToken(dataUser.Id)
	if err != nil {
		return response.ResponseLogin{}, err
	}

	refreshToken, err := user.auth.CreateRefreshToken(dataUser.Id)
	if err != nil {
		return response.ResponseLogin{}, err
	}

	response := response.ModelToResponseLogin(dataUser.Username, accesToken, refreshToken)
	return response, nil
}

func (user *userService) RefreshToken(refresh string) (response.ResponseLogin, error) {
	userId, err := user.auth.ParseToken(refresh, config.InitConfig().REFRESH_TOKEN)
	if err != nil {
		return response.ResponseLogin{}, err
	}

	accesToken, err := user.auth.CreateAccessToken(userId)
	if err != nil {
		return response.ResponseLogin{}, err
	}

	refreshToken, err := user.auth.CreateRefreshToken(userId)
	if err != nil {
		return response.ResponseLogin{}, err
	}

	response := response.ModelToResponseLogin("", accesToken, refreshToken)
	return response, nil
}

func (user *userService) SelectUserById(id string) (response.ResponseUser, error) {
	dataUser, err := user.userRepository.SelectUserById(id)
	if err != nil {
		return response.ResponseUser{}, err
	}

	return dataUser, nil
}
