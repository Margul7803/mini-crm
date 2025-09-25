package store

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// GORMStore implements the Storer interface using GORM.
// It stores contacts in a SQLite database.
type GORMStore struct {
	db *gorm.DB
}

// NewGORMStore creates a new GORMStore and initializes the database.
func NewGORMStore(dbPath string) (*GORMStore, error) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto-migrate the schema
	db.AutoMigrate(&Contact{})

	return &GORMStore{db: db}, nil
}

// Add adds a new contact to the database.
func (s *GORMStore) Add(contact Contact) error {
	return s.db.Create(&contact).Error
}

// Get retrieves a contact from the database by its ID.
func (s *GORMStore) Get(id uint) (Contact, error) {
	var contact Contact
	err := s.db.First(&contact, id).Error
	return contact, err
}

// GetAll retrieves all contacts from the database.
func (s *GORMStore) GetAll() ([]Contact, error) {
	var contacts []Contact
	err := s.db.Find(&contacts).Error
	return contacts, err
}

// Update updates an existing contact in the database.
func (s *GORMStore) Update(contact Contact) error {
	return s.db.Save(&contact).Error
}

// Delete removes a contact from the database by its ID.
func (s *GORMStore) Delete(id uint) error {
	return s.db.Delete(&Contact{}, id).Error
}