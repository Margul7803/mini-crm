package store

import (
	"encoding/json"
	"os"
	"sync"
)

// JSONStore implements the Storer interface using a JSON file.
// It stores contacts in a JSON file.
type JSONStore struct {
	mu   sync.RWMutex
	path string
}

// NewJSONStore creates a new JSONStore.
func NewJSONStore(path string) *JSONStore {
	return &JSONStore{path: path}
}

func (s *JSONStore) readContacts() (map[uint]Contact, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	data, err := os.ReadFile(s.path)
	if err != nil {
		if os.IsNotExist(err) {
			return make(map[uint]Contact), nil
		}
		return nil, err
	}

	var contacts map[uint]Contact
	if err := json.Unmarshal(data, &contacts); err != nil {
		return nil, err
	}
	return contacts, nil
}

func (s *JSONStore) writeContacts(contacts map[uint]Contact) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	data, err := json.MarshalIndent(contacts, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.path, data, 0644)
}

func (s *JSONStore) Add(contact Contact) error {
	contacts, err := s.readContacts()
	if err != nil {
		return err
	}

	var maxID uint
	for id := range contacts {
		if id > maxID {
			maxID = id
		}
	}
	contact.ID = maxID + 1

	contacts[contact.ID] = contact
	return s.writeContacts(contacts)
}

func (s *JSONStore) Get(id uint) (Contact, error) {
	contacts, err := s.readContacts()
	if err != nil {
		return Contact{}, err
	}

	contact, ok := contacts[id]
	if !ok {
		return Contact{}, os.ErrNotExist
	}
	return contact, nil
}

func (s *JSONStore) GetAll() ([]Contact, error) {
	contacts, err := s.readContacts()
	if err != nil {
		return nil, err
	}

	var result []Contact
	for _, contact := range contacts {
		result = append(result, contact)
	}
	return result, nil
}

func (s *JSONStore) Update(contact Contact) error {
	contacts, err := s.readContacts()
	if err != nil {
		return err
	}

	_, ok := contacts[contact.ID]
	if !ok {
		return os.ErrNotExist
	}

	contacts[contact.ID] = contact
	return s.writeContacts(contacts)
}

func (s *JSONStore) Delete(id uint) error {
	contacts, err := s.readContacts()
	if err != nil {
		return err
	}

	_, ok := contacts[id]
	if !ok {
		return os.ErrNotExist
	}

	delete(contacts, id)
	return s.writeContacts(contacts)
}
