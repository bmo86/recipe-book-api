package models

import "gorm.io/gorm"

type Recipe struct {
	gorm.Model  // add fields "id", "created_at", "updated_at", "deleted_at"
	Name        string
	Ingredients string
	Status      bool
	UserID      uint // ID_user create to recipe
}
