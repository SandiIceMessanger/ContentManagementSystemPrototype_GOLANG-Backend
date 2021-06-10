package database

import (
	"project/config"
	"project/middlewares"
	"project/models"
)

func GetUsers() (interface{}, error) {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUser(id int) (interface{}, error) {
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func DeleteUser(id int) (interface{}, error) {
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return nil, err
	}

	if deleteErr := config.DB.Delete(&user).Error; deleteErr != nil {
		return nil, deleteErr
	}
	return nil, nil
}

func UpdateUser(id int, user *models.User) (interface{}, error) {
	var existingUser models.User
	if err := config.DB.First(&existingUser, id).Error; err != nil {
		return nil, err
	}

	existingUser.Name = user.Name
	existingUser.Email = user.Email
	if updateErr := config.DB.Save(&existingUser).Error; updateErr != nil {
		return nil, updateErr
	}
	
	return existingUser, nil
}

func CreateUser(user *models.User) (interface{}, error) {
	if err := config.DB.Create(user).Error; err != nil {
		return nil, err
	}

	var err error
	user.Token, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	} 

	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func LoginUsers(user *models.User) (interface{}, error) {
	var err error
	if err = config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(user).Error; err != nil {
		return nil, err
	}

	user.Token, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}

	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}