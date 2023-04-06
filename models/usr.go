package models

import "gorm.io/gorm"

type User struct {
	gorm.Model   // add fields "id", "created_at", "updated_at", "deleted_at"
	Name         string
	Email        string
	AuthProvider string // "email" or "google"
	HashPass     string // hash of the pass (is auth email with pass)
	AuthToken    string // token of access (is auth with Google)
}
