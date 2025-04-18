# LRUCache (Least Recently Used Cache) TEST1

This package provides a Go implementation of a **Least Recently Used (LRU)** cache. It offers a simple yet efficient way to store data with quick access, automatically removing the least recently accessed items when the cache reaches its capacity.

## Features

- **O(1) Access and Insertion**: Provides constant-time operations for getting and putting items.
- **Doubly-Linked List and Map**: Uses an internal map for fast lookups and a doubly-linked list to track item usage.
- **Generic Value Storage**: Supports storing values of any type (`any`).

## Installation

To use this package, simply import it into your Go project:

```bash
go get github.com/michaelwp/lrucache
```

## Usage

### Creating a Cache

```go
cache := lrucache.NewLRUCache(3) // Capacity set to 3
```

### Adding Items

```go
cache.Put("key1", "value1")
cache.Put("key2", "value2")
cache.Put("key3", "value3")
```

### Retrieving Items

```go
value := cache.Get("key1") // Moves key1 to most recently used
fmt.Println(value) // Output: value1
```

### Cache Eviction

Adding more items beyond the cache capacity evicts the least recently used item:

```go
cache.Put("key4", "value4") // Evicts "key2" (if "key1" was recently accessed)
```

### Viewing Cache Content

To inspect the current content of the cache:

```go
list := cache.ViewList()
for _, item := range list {
    fmt.Println(item.Key, item.Value)
}
```

## Types and Functions

### Interface `LRUCache`

```go
type LRUCache interface {
	Get(key string) any
	Put(key string, value any)
	ViewList() []*data
}
```

### Functions

- `NewLRUCache(capacity int) LRUCache`: Creates a new cache with a specified capacity.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.
