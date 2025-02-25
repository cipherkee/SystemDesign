package main

import (
	"errors"
	"reflect"
	"sync"

	"github.com/cipherkee/SystemDesign/InmemoryKeyValueStore/value"
)

var (
	NoKeyError      = errors.New("no such key")
	UpdateIsInvalid = errors.New("Update is invalid")
)

type KeyValueStore struct {
	data map[string]*value.Value

	/* attribute name: attribute object.
	if all instances of the attribute are deleted also, the type will not change
	*/
	attributeValueType map[string]string // TODO: Ensure type is maintained across values
	// attr key: attr value: key
	searchIndex map[string]map[interface{}][]string // TODO: search to be indexed
	// Locking
	mu sync.RWMutex
}

func NewKeyValueStore() *KeyValueStore {
	return &KeyValueStore{
		data:               map[string]*value.Value{},
		attributeValueType: map[string]string{},
		searchIndex:        map[string]map[interface{}][]string{},
		mu:                 sync.RWMutex{},
	}
}

func (kv *KeyValueStore) Get(key string) (map[string]interface{}, error) {
	kv.mu.RLock()
	defer kv.mu.RUnlock()
	val, ok := kv.data[key]
	if !ok {
		return nil, NoKeyError
	}
	return val.GetValueAsMap()
}

func (kv *KeyValueStore) Put(key string, val map[string]interface{}) error {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	// Validate attribute types
	isValidUpdate := kv.validateAttributes(val)
	if !isValidUpdate {
		return UpdateIsInvalid
	} else {
		kv.updateAttributes(val)
	}

	kv.updateIndexWithNewKeyValue(key, val)

	currVal, exist := kv.data[key]
	if !exist {
		kv.data[key] = value.NewValue(val)
		return nil
	}
	return currVal.SetAttributes(val)
}

func (kv *KeyValueStore) Delete(key string) error {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	_, ok := kv.data[key]
	if !ok {
		return NoKeyError
	}
	delete(kv.data, key)
	return nil
}

func (kv *KeyValueStore) Keys() []string {
	kv.mu.RLock()
	defer kv.mu.RUnlock()
	keys := make([]string, 0)
	for key, _ := range kv.data {
		keys = append(keys, key)
	}
	return keys // TODO: Can optimize by indexing the keys
}

func (kv *KeyValueStore) search(attributeKey string, attributeValue string) []string {
	kv.mu.RLock()
	defer kv.mu.RUnlock()
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

func (kv *KeyValueStore) validateAttributes(attrMap map[string]interface{}) bool {
	for k, v := range attrMap {
		currValType, ok := kv.attributeValueType[k]
		if ok {
			if currValType != reflect.TypeOf(v).String() {
				return false
			}
		}
	}
	return true
}

func (kv *KeyValueStore) updateAttributes(attrMap map[string]interface{}) {
	for k, v := range attrMap {
		switch v.(type) {
		case string:
			kv.attributeValueType[k] = "string"
		case int:
			kv.attributeValueType[k] = "int"
		default:
			return
		}
	}
}

func (kv *KeyValueStore) updateIndexWithNewKeyValue(key string, attrMap map[string]interface{}) {
	for k, v := range attrMap {
		kv.singleAttrIndexing(key, k, v)
	}
	return
}

func (kv *KeyValueStore) singleAttrIndexing(key string, attrKey string, attrVal interface{}) {
	attrIndex, ok := kv.searchIndex[attrKey]
	if !ok {
		kv.searchIndex[attrKey] = map[interface{}][]string{}
		attrIndex = kv.searchIndex[attrKey]
	}
	_, ok = attrIndex[attrVal]
	if !ok {
		attrIndex[attrVal] = make([]string, 0)
	}
	attrIndex[attrVal] = append(attrIndex[attrVal], key)
	return
}

func (kv *KeyValueStore) SearchAttr(attrKey string, attrValue interface{}) []string {
	kv.mu.RLock()
	defer kv.mu.RUnlock()
	attrIndex, ok := kv.searchIndex[attrKey]
	if !ok {
		return nil
	}
	keys, ok := attrIndex[attrValue]
	if !ok {
		return nil
	}
	return keys
}
