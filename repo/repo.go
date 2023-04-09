package repo

import (
	"recipe-book-api/models"

	"golang.org/x/oauth2"
)

type RepoUser interface {
	SingUpNormal(usr *models.User) (uint, error)
	SingUpGoogle(usr *models.User) (oauth2.AuthCodeOption, error)
	GetUserById(id uint) (*models.User, error)
	GetUserEmail(email string) (*models.User, error)
}

var repo_user RepoUser

func SetRepoUser(repo RepoUser) {
	repo_user = repo
}

func SingUpNormal(usr *models.User) (uint, error) {
	return repo_user.SingUpNormal(usr)
}

func SingUpGoogle(usr *models.User) (oauth2.AuthCodeOption, error) {
	return repo_user.SingUpGoogle(usr)
}

func GetUserById(id uint) (*models.User, error) {
	return repo_user.GetUserById(id)
}

func GetUserEmail(email string) (*models.User, error) {
	return repo_user.GetUserEmail(email)
}
