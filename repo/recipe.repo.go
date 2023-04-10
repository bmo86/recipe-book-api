package repo

import "recipe-book-api/models"

// functions of recipe
func NewRecipe(recipe *models.Recipe) (uint, error)   { return repo_app.NewRecipe(recipe) }
func DelteRecipe(id uint) error                       { return repo_app.DelteRecipe(id) }
func UpdateRecipe(recipe models.Recipe) (uint, error) { return repo_app.UpdateRecipe(recipe) }
func GetRecipe(id uint) (*models.Recipe, error)       { return repo_app.GetRecipe(id) }
