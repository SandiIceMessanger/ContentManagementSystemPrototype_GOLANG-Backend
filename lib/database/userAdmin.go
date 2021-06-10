package database

import (
	"project/config"
	"project/middlewares"
	"project/models"
)

func GetUserAdmins() (interface{}, error) {
	var userAdmins []models.UserAdmin

	if err := config.DB.Find(&userAdmins).Error; err != nil {
		return nil, err
	}
	return userAdmins, nil
}

func GetUserAdmin(id int) (interface{}, error) {
	var userAdmin models.UserAdmin

	if err := config.DB.First(&userAdmin, id).Error; err != nil {
		return nil, err
	}
	return userAdmin, nil
}

func DeleteUserAdmin(id int) (interface{}, error) {
	var userAdmin models.UserAdmin
	if err := config.DB.First(&userAdmin, id).Error; err != nil {
		return nil, err
	}

	if deleteErr := config.DB.Delete(&userAdmin).Error; deleteErr != nil {
		return nil, deleteErr
	}
	return nil, nil
}

func UpdateUserAdmin(id int, userAdmin *models.UserAdmin) (interface{}, error) {
	var existingUserAdmin models.UserAdmin
	if err := config.DB.First(&existingUserAdmin, id).Error; err != nil {
		return nil, err
	}

	existingUserAdmin.Name = userAdmin.Name
	existingUserAdmin.Email = userAdmin.Email
	if updateErr := config.DB.Save(&existingUserAdmin).Error; updateErr != nil {
		return nil, updateErr
	}

	return existingUserAdmin, nil
}

func CreateUserAdmin(userAdmin *models.UserAdmin) (interface{}, error) {
	if err := config.DB.Create(userAdmin).Error; err != nil {
		return nil, err
	}

	var err error
	userAdmin.Token, err = middlewares.CreateToken(int(userAdmin.ID))
	if err != nil {
		return nil, err
	}

	if err := config.DB.Save(userAdmin).Error; err != nil {
		return nil, err
	}

	return userAdmin, nil
}

func LoginUserAdmins(userAdmin *models.UserAdmin) (interface{}, error) {
	var err error
	if err = config.DB.Where("email = ? AND password = ?", userAdmin.Email, userAdmin.Password).First(userAdmin).Error; err != nil {
		return nil, err
	}

	userAdmin.Token, err = middlewares.CreateToken(int(userAdmin.ID))
	if err != nil {
		return nil, err
	}

	if err := config.DB.Save(userAdmin).Error; err != nil {
		return nil, err
	}

	return userAdmin, nil
}
