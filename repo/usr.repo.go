package repo

import (
	"recipe-book-api/models"

	"golang.org/x/oauth2"
)

// Fucntions of users

func SingUpNormal(usr *models.User) (uint, error) { return repo_app.SingUpNormal(usr) }

func SingUpGoogle(usr *models.User) (oauth2.AuthCodeOption, error) { return repo_app.SingUpGoogle(usr) }

func GetUserById(id uint) (*models.User, error) { return repo_app.GetUserById(id) }

func GetUserEmail(email string) (*models.User, error) { return repo_app.GetUserEmail(email) }
