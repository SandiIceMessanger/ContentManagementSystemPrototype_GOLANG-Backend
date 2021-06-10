package database

import (
	"project/config"
	"project/middlewares"
	"project/models"
)

func GetUserClients() (interface{}, error) {
	var userClients []models.UserClient

	if err := config.DB.Find(&userClients).Error; err != nil {
		return nil, err
	}
	return userClients, nil
}

func GetUserClient(id int) (interface{}, error) {
	var userClient models.UserClient

	if err := config.DB.First(&userClient, id).Error; err != nil {
		return nil, err
	}
	return userClient, nil
}

func DeleteUserClient(id int) (interface{}, error) {
	var userClient models.UserClient
	if err := config.DB.First(&userClient, id).Error; err != nil {
		return nil, err
	}

	if deleteErr := config.DB.Delete(&userClient).Error; deleteErr != nil {
		return nil, deleteErr
	}
	return nil, nil
}

func UpdateUserClient(id int, userClient *models.UserClient) (interface{}, error) {
	var existingUserClient models.UserClient
	if err := config.DB.First(&existingUserClient, id).Error; err != nil {
		return nil, err
	}

	existingUserClient.Name = userClient.Name
	existingUserClient.Email = userClient.Email
	if updateErr := config.DB.Save(&existingUserClient).Error; updateErr != nil {
		return nil, updateErr
	}

	return existingUserClient, nil
}

func CreateUserClient(userClient *models.UserClient) (interface{}, error) {
	if err := config.DB.Create(userClient).Error; err != nil {
		return nil, err
	}

	var err error
	userClient.Token, err = middlewares.CreateToken(int(userClient.ID))
	if err != nil {
		return nil, err
	}

	if err := config.DB.Save(userClient).Error; err != nil {
		return nil, err
	}

	return userClient, nil
}

func LoginUserClients(userClient *models.UserClient) (interface{}, error) {
	var err error
	if err = config.DB.Where("email = ? AND password = ?", userClient.Email, userClient.Password).First(userClient).Error; err != nil {
		return nil, err
	}

	userClient.Token, err = middlewares.CreateToken(int(userClient.ID))
	if err != nil {
		return nil, err
	}

	if err := config.DB.Save(userClient).Error; err != nil {
		return nil, err
	}

	return userClient, nil
}
