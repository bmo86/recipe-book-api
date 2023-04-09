package repo

import "recipe-book-api/models"

type RepoRecipe interface {
	NewRecipe(recipe *models.Recipe) (uint, error)
	DelteRecipe(id uint) error
	UpdateRecipe(recipe models.Recipe) (uint, error)
	GetRecipe(id uint) (*models.Recipe, error)
}

var repo_recipe RepoRecipe

func SetReporecipe(r RepoRecipe) {
	repo_recipe = r
}

func NewRecipe(recipe *models.Recipe) (uint, error)   { return repo_recipe.NewRecipe(recipe) }
func DelteRecipe(id uint) error                       { return repo_recipe.DelteRecipe(id) }
func UpdateRecipe(recipe models.Recipe) (uint, error) { return repo_recipe.UpdateRecipe(recipe) }
func GetRecipe(id uint) (*models.Recipe, error)       { return repo_recipe.GetRecipe(id) }
