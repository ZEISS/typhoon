package storage

import (
	"sync"
)

// Storage is a simple object that provides thread safe
// methods to read and write into a map.
type Storage struct {
	data map[string]map[string]interface{}
	mux  sync.RWMutex
}

// New returns an instance of Storage.
func New() *Storage {
	return &Storage{
		data: make(map[string]map[string]interface{}),
		mux:  sync.RWMutex{},
	}
}

// Set writes a value interface to a string key.
func (s *Storage) Set(eventID, key string, value interface{}) {
	s.mux.Lock()
	defer s.mux.Unlock()
	if s.data[eventID] == nil {
		s.data[eventID] = make(map[string]interface{})
	}
	s.data[eventID][key] = value
}

// Get reads value by a key.
func (s *Storage) Get(eventID string, key string) interface{} {
	s.mux.RLock()
	defer s.mux.RUnlock()
	if s.data[eventID] == nil {
		return nil
	}
	return s.data[eventID][key]
}

// ListEventVariables returns the slice of variables created for EventID.
func (s *Storage) ListEventVariables(eventID string) []string {
	s.mux.RLock()
	defer s.mux.RUnlock()
	list := []string{}
	for k := range s.data[eventID] {
		list = append(list, k)
	}
	return list
}

// ListEventIDs returns the list of stored event IDs.
func (s *Storage) ListEventIDs() []string {
	s.mux.RLock()
	defer s.mux.RUnlock()
	list := []string{}
	for k := range s.data {
		list = append(list, k)
	}
	return list
}

// Flush removes variables by their parent event ID.
func (s *Storage) Flush(eventID string) {
	s.mux.Lock()
	defer s.mux.Unlock()
	delete(s.data, eventID)
}
