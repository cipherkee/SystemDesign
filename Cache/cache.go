package main

import (
	"errors"
)

type ICache interface {
	Get(key string) (any, error)
	Put(key string) (any, error)
	Delete(key string) (any, error)
}

type CacheLRU struct {
	size    int
	ep      IEvictionPolicy
	storage IStorage
}

func NewCacheLRU() *CacheLRU {
	ep := NewLRU()
	storage := NewStorageHashMap()
	return &CacheLRU{
		ep:      ep,
		storage: storage,
		size:    3, // TODO: Should this be here ? or in storage ?
	}
}

func (c *CacheLRU) Get(key string) (any, error) {
	if err := c.ep.Update(key); err != nil {
		return nil, err
	}

	return c.storage.Get(key)
}

func (c *CacheLRU) Put(key string, val any) error {
	currSize := c.storage.Size()
	// eviction if storage is full
	if currSize == c.size {
		evictKey, err := c.ep.GetEvictionItem()
		if err != nil {
			return errors.New("unable to evict key")
		}
		err = c.ep.Delete(evictKey)
		if err != nil {
			return errors.New("unable to delete key")
		}
		c.storage.Delete(evictKey)
	}

	c.storage.Put(key, val)
	c.ep.Update(key)
	return nil
}

func (c *CacheLRU) Delete(key string) error {
	err := c.ep.Delete(key)
	if err != nil {
		return err
	}
	return c.storage.Delete(key)
}
