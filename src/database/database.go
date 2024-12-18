package database

type DB[K, V any] interface {
	Get(key K) (V, error)
	Set(key K, value V) error
	Delete(key K) error
	Init() error
	Close() error
}