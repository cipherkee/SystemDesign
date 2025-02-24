package main

import (
	"fmt"
)

/*
The key will always be a string.
The value would be an object/map. The object would have attributes and corresponding values.
Each attribute key would be a string and the attribute values could be string, integer, double or boolean.

The key-value store should be thread-safe.
1. get(String key) => Should return the value (object with attributes and their values). Return null if key not present
2. search(String attributeKey, String attributeValue) => Returns a list of keys that have the given attribute key, value pair.
3. put(String key, List<Pair<String, String>> listOfAttributePairs) => Adds the key and the attributes to the key-value store. If the key already exists then the value is replaced.
4. delete(String key) => Deletes the key, value pair from the store.
5. keys() => Return a list of all the keys

The value object should override the toString
The data type of an attribute should get fixed after its first occurrence.
Nothing should be printed inside any of these methods. All scanning and printing should happen in the Driver/Main class only. Exception Handling should also happen in the Driver/Main class.
*/

func main() {
	kvstore := NewKeyValueStore()

	kvstore.Put("key1", map[string]interface{}{
		"name":   "Pomogranate",
		"colour": "pink",
		"count":  3,
	})

	kvstore.Put("key2", map[string]interface{}{
		"name":   "Apple",
		"colour": "red",
		"count":  2,
	})

	s, err := kvstore.Get("key2")
	if err != nil {
		panic(err)
	}
	fmt.Println(s)

	keys := kvstore.Keys()
	fmt.Println(keys)

	kvstore.Delete("key1")
	keys = kvstore.Keys()
	fmt.Println(keys)
}
