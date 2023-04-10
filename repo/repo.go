package repo

import (
	"recipe-book-api/models"

	"golang.org/x/oauth2"
)

type RepoApp interface {
	// functions usr
	SingUpNormal(usr *models.User) (uint, error)
	SingUpGoogle(usr *models.User) (oauth2.AuthCodeOption, error)
	GetUserById(id uint) (*models.User, error)
	GetUserEmail(email string) (*models.User, error)

	//functions recipe
	NewRecipe(recipe *models.Recipe) (uint, error)
	DelteRecipe(id uint) error
	UpdateRecipe(recipe models.Recipe) (uint, error)
	GetRecipe(id uint) (*models.Recipe, error)
}

var repo_app RepoApp

func SetRepo(repo RepoApp) {
	repo_app = repo
}
