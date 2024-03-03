package repository

import (
	"refresh/internal/user/dto/request"
	"refresh/internal/user/dto/response"
	"refresh/internal/user/model"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepositoryInterface interface {
	InsertUser(data request.UserRequest) error
	SelectUsername(username string) (model.Users, error)
	SelectUserById(id string) (response.ResponseUser, error)
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &userRepository{
		db: db,
	}
}

func (user *userRepository) InsertUser(data request.UserRequest) error {
	request := request.UserRequestToModel(data)

	tx := user.db.Create(&request)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (user *userRepository) SelectUsername(username string) (model.Users, error) {
	dataUser := model.Users{}

	tx := user.db.Where("username = ?", username).Take(&dataUser)
	if tx.Error != nil {
		return model.Users{}, tx.Error
	}

	return dataUser, nil
}

func (user *userRepository) SelectUserById(id string) (response.ResponseUser, error) {
	dataUser := model.Users{}

	tx := user.db.Where("id = ?", id).First(&dataUser)
	if tx.Error != nil {
		return response.ResponseUser{}, tx.Error
	}

	response := response.ModelToResponseUser(dataUser)
	return response, nil
}