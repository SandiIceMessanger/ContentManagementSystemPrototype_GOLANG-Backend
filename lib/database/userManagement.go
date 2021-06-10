package database

import (
	"project/config"
	"project/middlewares"
	"project/models"
)

func GetUserManagements() (interface{}, error) {
	var userManagements []models.UserManagement

	if err := config.DB.Find(&userManagements).Error; err != nil {
		return nil, err
	}
	return userManagements, nil
}

func GetUserManagement(id int) (interface{}, error) {
	var userManagement models.UserManagement

	if err := config.DB.First(&userManagement, id).Error; err != nil {
		return nil, err
	}
	return userManagement, nil
}

func DeleteUserManagement(id int) (interface{}, error) {
	var userManagement models.UserManagement
	if err := config.DB.First(&userManagement, id).Error; err != nil {
		return nil, err
	}

	if deleteErr := config.DB.Delete(&userManagement).Error; deleteErr != nil {
		return nil, deleteErr
	}
	return nil, nil
}

func UpdateUserManagement(id int, userManagement *models.UserManagement) (interface{}, error) {
	var existingUserManagement models.UserManagement
	if err := config.DB.First(&existingUserManagement, id).Error; err != nil {
		return nil, err
	}

	existingUserManagement.Name = userManagement.Name
	existingUserManagement.Email = userManagement.Email
	if updateErr := config.DB.Save(&existingUserManagement).Error; updateErr != nil {
		return nil, updateErr
	}

	return existingUserManagement, nil
}

func CreateUserManagement(userManagement *models.UserManagement) (interface{}, error) {
	if err := config.DB.Create(userManagement).Error; err != nil {
		return nil, err
	}

	var err error
	userManagement.Token, err = middlewares.CreateToken(int(userManagement.ID))
	if err != nil {
		return nil, err
	}

	if err := config.DB.Save(userManagement).Error; err != nil {
		return nil, err
	}

	return userManagement, nil
}

func LoginUserManagements(userManagement *models.UserManagement) (interface{}, error) {
	var err error
	if err = config.DB.Where("email = ? AND password = ?", userManagement.Email, userManagement.Password).First(userManagement).Error; err != nil {
		return nil, err
	}

	userManagement.Token, err = middlewares.CreateToken(int(userManagement.ID))
	if err != nil {
		return nil, err
	}

	if err := config.DB.Save(userManagement).Error; err != nil {
		return nil, err
	}

	return userManagement, nil
}
