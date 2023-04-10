package database

import (
	"recipe-book-api/models"
)

func (i *instancePostgres) SingUpNormal(usr *models.User) (uint, error) {
	err := i.db.Create(usr)
	return usr.ID, err.Error
}

func (i *instancePostgres) GetUserById(id uint) (*models.User, error) {
	var user models.User
	err := i.db.Table("users").Select("name, email, authProvider, hashPass, authToken").First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (i *instancePostgres) GetUserEmail(email string) (*models.User, error) {
	var user models.User
	err := i.db.Table("users").Select("name, email, authProvider, hashPass, authToken").First(&user, email).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
