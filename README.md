# Rate limiter

Configurable rate limiting lib with multiple algorithms to choose from.

## Usage

```go
package main

import (
	"time"
	"github.com/valsov/rlimit/storage"
)

func main() {
	// Instanciate a new bucket limiter with a backing memory store
	memoryStore := storage.NewMemoryStore[BucketConfig, BucketValue]()
	config := BucketConfig{
		RefillRate: time.Hour, // Every hour
		RefillCount: 10, // Refill 10 tokens each RefillRate
		BucketSize: 15, // Bucket capacity
	}
	limiter := NewBucketLimiter(memoryStore, config)
	
	// Try to accept a request
	success, err := limiter.TryAllow("sample-id", 1)
	if err != nil {
		// Handle error
	}
	
	if success {
		// Request accepted
	} else {
		// Request rejected
	}
}
```

## Rate limiting algorithms
### Bucket

The token bucket algorithm regulates traffic flow by maintaining a bucket of tokens, each representing a unit of resource allowance. Requests consume tokens from the bucket, which refills at a set rate. If the bucket is empty, requests are rejected.

#### Data structures

```go
// Configuration of the limiter
type BucketConfig struct {
	RefillRate time.Duration
	RefillCount int // How much tokens should be periodically added to the bucket
	BucketSize int
}

// Stored value per user
type BucketValue struct {
	LastRefillUtc time.Time
	RemainingTokens int
}
```

### Leaky bucket

The leaky bucket algorithm controls traffic by using a bucket that leaks tokens at a fixed rate. Incoming requests are added to the bucket, but if it overflows, excess requests are rejected.

#### Data structures

```go
// Configuration of the limiter
type LeakyBucketConfig struct {
	NewTokensRate time.Duration
	NewTokensCount int
	BucketSize int
}

// Stored value per user
type LeakyBucketValue struct {
	LastCheckedUtc time.Time
	RemainingTokens int
}
```

### Fixed window

The fixed window algorithm tracks the number of requests within predefined time windows. Each window has a fixed duration and number of allowed requests for this duration. Requests exceeding  this limit are delayed. Once the duration is over, the requests count is reset and the duration can start again.

#### Data structures

```go
// Configuration of the limiter
type FixedWindowConfig struct {
	WindowLength time.Duration
	MaxTokens int
}

// Stored value per user
type FixedWindowValue struct {
	WindowStartUtc time.Time
	RemainingTokens int
}
```

### Sliding window

The sliding window algorithm tracks requests within a moving time window, unlike the fixed window's static intervals. Requests are counted within a window that continuously slides forward in time. It offers more granular control over rate limiting compared to the fixed window method.

#### Data structures

```go
// Configuration of the limiter
type SlidingWindowConfig struct {
	WindowLength time.Duration
	MaxTokens int
}

// Stored value per user
type SlidingWindowValue struct {
	PreviousConsumption int
	CurrentConsumption int
	WindowStartUtc time.Time
}
```
