package database

import (
	"project/config"
	"project/middlewares"
	"project/models"
)

func GetPasswordForgets() (interface{}, error) {
	var passwordForgets []models.PasswordForget

	if err := config.DB.Find(&passwordForgets).Error; err != nil {
		return nil, err
	}
	return passwordForgets, nil
}

func GetPasswordForget(id int) (interface{}, error) {
	var passwordForget models.PasswordForget

	if err := config.DB.First(&passwordForget, id).Error; err != nil {
		return nil, err
	}
	return passwordForget, nil
}

func DeletePasswordForget(id int) (interface{}, error) {
	var passwordForget models.PasswordForget
	if err := config.DB.First(&passwordForget, id).Error; err != nil {
		return nil, err
	}

	if deleteErr := config.DB.Delete(&passwordForget).Error; deleteErr != nil {
		return nil, deleteErr
	}
	return nil, nil
}

func UpdatePasswordForget(id int, passwordForget *models.PasswordForget) (interface{}, error) {
	var existingPasswordForget models.PasswordForget
	if err := config.DB.First(&existingPasswordForget, id).Error; err != nil {
		return nil, err
	}

	existingPasswordForget.Token = passwordForget.Token
	existingPasswordForget.IdUser = passwordForget.IdUser

	if updateErr := config.DB.Save(&existingPasswordForget).Error; updateErr != nil {
		return nil, updateErr
	}

	return existingPasswordForget, nil
}

func CreatePasswordForget(passwordForget *models.PasswordForget) (interface{}, error) {
	if err := config.DB.Create(passwordForget).Error; err != nil {
		return nil, err
	}

	var err error
	passwordForget.Token, err = middlewares.CreateToken(int(passwordForget.ID))
	if err != nil {
		return nil, err
	}

	if err := config.DB.Save(passwordForget).Error; err != nil {
		return nil, err
	}

	return passwordForget, nil
}
