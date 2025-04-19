package main

import (
	"errors"
)

type IEvictionPolicy interface {
	GetEvictionItem() (string, error)
	Delete(string) error
	Update(string) error
}

type lruItem struct {
	key   string
	right *lruItem
	left  *lruItem
}

type LRU struct {
	head   *lruItem
	tail   *lruItem
	lookup map[string]*lruItem
}

func (l *LRU) GetEvictionItem() (string, error) {
	if l.tail == nil {
		return "", errors.New("empty cache, nothing to delete")
	}
	return l.tail.key, nil
}

func (l *LRU) Delete(key string) error {
	item, ok := l.lookup[key]
	if !ok {
		return errors.New("Invalid key")
	}

	delete(l.lookup, key)
	left := item.left
	right := item.right

	if left == nil {
		l.head = right
		right.left = nil
		return nil
	}

	if right == nil {
		l.tail = left
		left.right = nil
		return nil
	}

	left.right = right
	right.left = left
	return nil
}

func (l *LRU) Update(key string) error {
	item, ok := l.lookup[key]
	if ok {
		left := item.left
		right := item.right

		if left == nil { // this key is already the most recently used
			return nil
		}
		left.right = right
		if right != nil {
			right.left = left
		}

	} else {
		item = &lruItem{
			key: key,
		}
		l.lookup[key] = item
	}

	head := l.head
	if head == nil {
		l.head = item
		l.tail = item
		return nil
	}
	item.right = head
	head.left = item
	l.head = item
	return nil
}

func NewLRU() *LRU {
	return &LRU{
		lookup: map[string]*lruItem{},
	}
}
