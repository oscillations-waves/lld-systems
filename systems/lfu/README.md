# LFU Cache System

This project implements an **LFU (Least Frequently Used) Cache** in Go.

## Overview

An LFU cache is a data structure that evicts the least frequently used items first when it reaches its capacity. If multiple items have the same frequency, the least recently used among them is evicted.

## Features
- O(1) time complexity for both `Get` and `Put` operations
- Frequency-based eviction policy
- Handles cache capacity and updates frequencies on access

## Usage

### Example
```go
cache := NewLFUCache(2)
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
- Each cache entry tracks its frequency and its position in a frequency list.
- The cache maintains a mapping from frequency to a doubly linked list of nodes.
- On `Get` or `Put`, the frequency of the node is increased and it is moved to the appropriate list.
- When the cache is full, the least frequently (and least recently) used node is evicted.

## Running the Code

1. Make sure you have Go installed.
2. Navigate to the `lfu-cache-system` directory.
3. Run:
   ```sh
   go run lfu_cache.go
   ```

## File Structure
- `lfu_cache.go`: Main implementation and example usage.

## License
MIT
