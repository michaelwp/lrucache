package main

import (
	"fmt"
	"github.com/michaelwp/lrucache"
)

func main() {
	lru := lrucache.NewLRUCache(2)

	lru.Put("one", 10)
	lru.Put("two", 20)
	fmt.Println("one", ":", lru.Get("one")) // Output: 10
	lru.Put("three", 30)                    // Evicts key "two"
	fmt.Println("two", ":", lru.Get("two")) // Output: nil (not found)

	fmt.Println("\n---key:value---")
	list1 := lru.ViewList()
	for _, data := range list1 {
		fmt.Printf("%s:%v\n", data.Key, data.Value)
	}

	fmt.Println("")
	lru.Put("four", 40)                     // Evicts key "one"
	fmt.Println("one", ":", lru.Get("one")) // Output: nil (not found)

	fmt.Println("\n---key:value---")
	list2 := lru.ViewList()
	for _, data := range list2 {
		fmt.Printf("%s:%v\n", data.Key, data.Value)
	}
}
