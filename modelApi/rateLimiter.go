package modelApi

import (
	"math"
	"sync"
	"time"
)

type RateLimiter struct {
	rate        float64
	burst       int
	bucketMutex sync.Mutex
	bucket      int
	lastRefill  time.Time
}

func NewRateLimiter(rate float64, burst int) *RateLimiter {
	return &RateLimiter{rate: rate, burst: burst, lastRefill: time.Now()}
}

func (rl *RateLimiter) Allow() bool {
	rl.bucketMutex.Lock()
	defer rl.bucketMutex.Unlock()

	now := time.Now()
	timeSinceLastRefill := now.Sub(rl.lastRefill).Seconds()
	tokensToAdd := float64(timeSinceLastRefill) * rl.rate
	rl.bucket = int(math.Min(float64(rl.burst), float64(rl.bucket)+tokensToAdd))
	rl.lastRefill = now

	if rl.bucket > 0 {
		rl.bucket--
		return true
	}

	return false
}
