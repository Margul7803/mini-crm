package store

import "gorm.io/gorm"

// Contact represents a contact in the CRM.
// Note: We redefine it here to avoid circular dependencies.
type Contact struct {
	gorm.Model
	Nom   string
	Email string
}

// Storer defines the interface for contact storage.
// This allows us to easily swap storage implementations.
type Storer interface {
	Add(contact Contact) error
	Get(id uint) (Contact, error)
	GetAll() ([]Contact, error)
	Update(contact Contact) error
	Delete(id uint) error
}
