package database

import (
	"project/config"
	"project/models"
)

func GetPermissionTransactions() (interface{}, error) {
	var permissionTransactions []models.PermissionTransaction

	if err := config.DB.Find(&permissionTransactions).Error; err != nil {
		return nil, err
	}
	return permissionTransactions, nil
}

func GetPermissionTransaction(id int) (interface{}, error) {
	var permissionTransaction models.PermissionTransaction

	if err := config.DB.First(&permissionTransaction, id).Error; err != nil {
		return nil, err
	}
	return permissionTransaction, nil
}

func DeletePermissionTransaction(id int) (interface{}, error) {
	var permissionTransaction models.PermissionTransaction
	if err := config.DB.First(&permissionTransaction, id).Error; err != nil {
		return nil, err
	}

	if deleteErr := config.DB.Delete(&permissionTransaction).Error; deleteErr != nil {
		return nil, deleteErr
	}
	return nil, nil
}

func UpdatePermissionTransaction(id int, permissionTransaction *models.PermissionTransaction) (interface{}, error) {
	var existingPermissionTransaction models.PermissionTransaction
	if err := config.DB.First(&existingPermissionTransaction, id).Error; err != nil {
		return nil, err
	}

	existingPermissionTransaction.Point = permissionTransaction.Point
	existingPermissionTransaction.IdPermissionMaster = permissionTransaction.IdPermissionMaster

	if updateErr := config.DB.Save(&existingPermissionTransaction).Error; updateErr != nil {
		return nil, updateErr
	}

	return existingPermissionTransaction, nil
}

func CreatePermissionTransaction(permissionTransaction *models.PermissionTransaction) (interface{}, error) {
	if err := config.DB.Create(permissionTransaction).Error; err != nil {
		return nil, err
	}

	if err := config.DB.Save(permissionTransaction).Error; err != nil {
		return nil, err
	}

	return permissionTransaction, nil
}
