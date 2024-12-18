package database

import (
	"errors"
)

// mockDB struct implements the DB interface
type MockDB[K comparable, V any] struct {
	data map[K]V
}

func (m *MockDB[K, V]) Get(key K) (V, error) {
	if val, ok := m.data[key]; ok {
		return val, nil
	}
	var zero V
	return zero, errors.New("Key not found")
}

func (m *MockDB[K, V]) Set(key K, value V) error {
	m.data[key] = value
	return nil
}

func (m *MockDB[K, V]) Delete(key K) error {
	if _, ok := m.data[key]; ok {
		delete(m.data, key)
		return nil
	}
	return errors.New("Key not found")
}

func (m *MockDB[K, V]) Init() error {
	m.data = make(map[K]V)
	return nil
}

func (m *MockDB[K, V]) Close() error {
	m.data = nil
	return nil
}