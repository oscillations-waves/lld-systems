# LRU Cache System

This project implements an **LRU (Least Recently Used) Cache** in Go.

## Overview

An LRU cache is a data structure that evicts the least recently used items first when it reaches its capacity. It is commonly used to manage memory or storage efficiently.

## Features
- O(1) time complexity for both `Get` and `Put` operations
- Recency-based eviction policy
- Uses Go's `container/list` for efficient queue management

## Usage

### Example
```go
cache := NewLRUCache(2)
cache.Put(1, 10)
cache.Put(2, 20)
fmt.Println(cache.Get(1)) // Output: 10
cache.Put(3, 30)          // Evicts key 2
fmt.Println(cache.Get(2)) // Output: -1
cache.Put(4, 40)          // Evicts key 1
fmt.Println(cache.Get(1)) // Output: -1
fmt.Println(cache.Get(3)) // Output: 30
fmt.Println(cache.Get(4)) // Output: 40
```

## How It Works
- Each cache entry is stored in a map for O(1) access.
- A doubly linked list tracks the usage order of keys.
- On `Get` or `Put`, the accessed/updated key is moved to the front (most recently used).
- When the cache is full, the least recently used key (at the back) is evicted.

## Running the Code

1. Make sure you have Go installed.
2. Navigate to the `lru-cache-system` directory.
3. Run:
   ```sh
   go run lru_cache.go
   ```

## File Structure
- `lru_cache.go`: Main implementation and example usage.

## License
MIT
