package store

import ("sync")

type Store struct {
    data map[string]Item
    mu   sync.RWMutex
}


	func NewStore() *Store {
	return &Store{
		data: make(map[string]Item),
	}
}


func (s *Store) Set(key string, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = Item{Value: value}

}

func (s *Store) Get(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	item, exists := s.data[key]
	if !exists {
		return "", false
	}
	return item.Value, true
}


func (s *Store) Delete(key string) {

	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, key)
}

func (s *Store) Exists(key string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	_, exists := s.data[key]
	return exists
}