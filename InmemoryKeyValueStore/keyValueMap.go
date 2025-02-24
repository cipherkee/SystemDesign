package main

import (
	"errors"
	"sync"

	"github.com/cipherkee/SystemDesign/InmemoryKeyValueStore/value"
)

var (
	NoKeyError = errors.New("no such key")
)

type KeyValueStore struct {
	data sync.Map // map[string]*value.Value

	attributeValueType map[string]value.Attribute // TODO: Ensure type is maintained across values
	// attr key: attr value: key
	searchIndex map[string]map[interface{}][]string // TODO: search to be indexed
}

func NewKeyValueStore() *KeyValueStore {
	return &KeyValueStore{
		data: sync.Map{},
	}
}

func (kv *KeyValueStore) Get(key string) (map[string]interface{}, error) {
	val, ok := kv.data.Load(key)
	if !ok {
		return nil, NoKeyError
	}
	return val.(*value.Value).GetValueAsMap()
}

func (kv *KeyValueStore) Put(key string, val map[string]interface{}) error {
	currVal, exist := kv.data.Load(key)
	if !exist {
		kv.data.Store(key, value.NewValue(val))
		return nil
	}
	return currVal.(*value.Value).SetAttributes(val)
}

func (kv *KeyValueStore) Delete(key string) error {
	_, ok := kv.data.Load(key)
	if !ok {
		return NoKeyError
	}
	kv.data.Delete(key)
	return nil
}

func (kv *KeyValueStore) Keys() []string {
	keys := make([]string, 0)
	kv.data.Range(func(key, value any) bool {
		keys = append(keys, key.(string))
		return true
	})
	return keys // TODO: Can optimize by indexing the keys
}

func (kv *KeyValueStore) search(attributeKey string, attributeValue string) []string {
	attrValMap, ok := kv.searchIndex[attributeKey]
	if !ok {
		return nil
	}
	keys, ok := attrValMap[attrValMap]
	if !ok {
		return nil
	}
	return keys
}

func (kv *KeyValueStore) UpdateIndexWithNewKeyValue(val *value.Value) {

}
