package store

import (
	"fmt"
	"sync"
)

// MemoryStore implements the Storer interface using an in-memory map.
type MemoryStore struct {
	mu       sync.RWMutex
	contacts map[uint]Contact
	nextID   uint
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		contacts: make(map[uint]Contact),
		nextID:   1,
	}
}

func (s *MemoryStore) Add(contact Contact) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	contact.ID = s.nextID
	s.contacts[s.nextID] = contact
	s.nextID++
	return nil
}

func (s *MemoryStore) Get(id uint) (Contact, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	contact, ok := s.contacts[id]
	if !ok {
		return Contact{}, fmt.Errorf("contact with id %d not found", id)
	}
	return contact, nil
}

func (s *MemoryStore) GetAll() ([]Contact, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var contacts []Contact
	for _, contact := range s.contacts {
		contacts = append(contacts, contact)
	}
	return contacts, nil
}

func (s *MemoryStore) Update(contact Contact) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.contacts[contact.ID]
	if !ok {
		return fmt.Errorf("contact with id %d not found", contact.ID)
	}

	s.contacts[contact.ID] = contact
	return nil
}

func (s *MemoryStore) Delete(id uint) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.contacts[id]
	if !ok {
		return fmt.Errorf("contact with id %d not found", id)
	}

	delete(s.contacts, id)
	return nil
}
