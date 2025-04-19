package main

import (
	"errors"
)

type IStorage interface {
	Get(string) (any, error)
	Put(string, any)
	Delete(string) error
	Size() int
}

type StorageHashMap struct {
	store map[string]any
}

func NewStorageHashMap() *StorageHashMap {
	return &StorageHashMap{
		store: map[string]any{},
	}
}

func (s *StorageHashMap) Get(key string) (any, error) {
	val, ok := s.store[key]
	if !ok {
		return nil, errors.New("Key not found")
	}
	return val, nil
}

func (s *StorageHashMap) Put(key string, val any) {
	s.store[key] = val
}

func (s *StorageHashMap) Delete(key string) error {
	delete(s.store, key)
	return nil
}

func (s *StorageHashMap) Size() int {
	return len(s.store)
}
