package store

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type GORMStore struct {
	db *gorm.DB
}

func NewGORMStore(dbPath string) (*GORMStore, error) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Contact{})

	return &GORMStore{db: db}, nil
}

func (s *GORMStore) Add(contact Contact) error {
	return s.db.Create(&contact).Error
}

func (s *GORMStore) Get(id uint) (Contact, error) {
	var contact Contact
	err := s.db.First(&contact, id).Error
	return contact, err
}

func (s *GORMStore) GetAll() ([]Contact, error) {
	var contacts []Contact
	err := s.db.Find(&contacts).Error
	return contacts, err
}

func (s *GORMStore) Update(contact Contact) error {
	return s.db.Save(&contact).Error
}

func (s *GORMStore) Delete(id uint) error {
	return s.db.Delete(&Contact{}, id).Error
}
