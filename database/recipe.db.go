package database

import "recipe-book-api/models"

func (i *instancePostgres) NewRecipe(recipe *models.Recipe) (uint, error) {
	err := i.db.Create(recipe)
	return recipe.ID, err.Error
}

func (i *instancePostgres) DelteRecipe(id uint) error {
	var recipe *models.Recipe
	if err := i.db.Delete(&recipe, id).Error; err != nil {
		return err
	}
	return nil
}

func (i *instancePostgres) UpdateRecipe(recipe models.Recipe) (uint, error) {
	data := map[string]interface{}{
		"name":        recipe.Name,
		"ingredients": recipe.Ingredients,
		"status":      recipe.Status,
		"updated_at":  recipe.UpdatedAt,
	}

	re := i.db.Table("recipes").Where("id = ?", recipe.ID).UpdateColumns(data)
	if re.Error != nil {
		return 0, re.Error
	}

	return recipe.ID, nil
}

func (i *instancePostgres) GetRecipe(id uint) (*models.Recipe, error) {
	var recipe models.Recipe
	err := i.db.First(&recipe, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &recipe, nil
}
