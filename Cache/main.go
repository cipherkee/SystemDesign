package main

import (
	"fmt"
)

func main() {
	cache := NewCacheLRU()

	cache.Put("key1", "val1")
	val, err := cache.Get("key1")
	if err != nil {
		panic(err)
	}
	fmt.Println(val)

	cache.Put("key2", "val2")
	cache.Put("key3", "val3")
	cache.Put("key4", "val4")

	val, err = cache.Get("key1")
	if err != nil {
		fmt.Println(err)
	}
	val, _ = cache.Get("key2")
	fmt.Printf("key2 value is: %v\n", val)
	val, _ = cache.Get("key3")
	fmt.Printf("key3 value is: %v\n", val)
	val, _ = cache.Get("key4")
	fmt.Printf("key4 value is: %v\n", val)
}
