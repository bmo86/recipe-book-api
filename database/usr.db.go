package database

import (
	"recipe-book-api/models"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func (i *instancePostgres) SingUpNormal(usr *models.User) (uint, error) {
	err := i.db.Create(usr)
	return usr.ID, err.Error
}

func (i *instancePostgres) SingUpGoogle(usr *models.User) (oauth2.AuthCodeOption, error) {
	var credentials string
	cfg, err := google.ConfigFromJSON([]byte(credentials), "https://www.googleapis.com/auth/userinfo.email")
	if err != nil {
		return nil, err
	}
	cfg.RedirectURL = "http://localhost:8080/oauth2callback"
	authURL := cfg.AuthCodeURL("state-token", oauth2.AccessTypeOffline)

	return authURL, nil
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
