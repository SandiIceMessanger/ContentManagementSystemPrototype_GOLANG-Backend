package database

import (
	"project/config"
	"project/models"
)

func GetPermissionMasters() (interface{}, error) {
	var permissionMasters []models.PermissionMaster

	if err := config.DB.Find(&permissionMasters).Error; err != nil {
		return nil, err
	}
	return permissionMasters, nil
}

func GetPermissionMaster(id int) (interface{}, error) {
	var permissionMaster models.PermissionMaster

	if err := config.DB.First(&permissionMaster, id).Error; err != nil {
		return nil, err
	}
	return permissionMaster, nil
}

func DeletePermissionMaster(id int) (interface{}, error) {
	var permissionMaster models.PermissionMaster
	if err := config.DB.First(&permissionMaster, id).Error; err != nil {
		return nil, err
	}

	if deleteErr := config.DB.Delete(&permissionMaster).Error; deleteErr != nil {
		return nil, deleteErr
	}
	return nil, nil
}

func UpdatePermissionMaster(id int, permissionMaster *models.PermissionMaster) (interface{}, error) {
	var existingPermissionMaster models.PermissionMaster
	if err := config.DB.First(&existingPermissionMaster, id).Error; err != nil {
		return nil, err
	}

	existingPermissionMaster.Permission = permissionMaster.Permission
	existingPermissionMaster.IdUser = permissionMaster.IdUser

	if updateErr := config.DB.Save(&existingPermissionMaster).Error; updateErr != nil {
		return nil, updateErr
	}

	return existingPermissionMaster, nil
}

func CreatePermissionMaster(permissionMaster *models.PermissionMaster) (interface{}, error) {
	if err := config.DB.Create(permissionMaster).Error; err != nil {
		return nil, err
	}

	if err := config.DB.Save(permissionMaster).Error; err != nil {
		return nil, err
	}

	return permissionMaster, nil
}
