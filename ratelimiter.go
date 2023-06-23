package hqgolimit

import (
	"sync"
	"time"
)

// RateLimiter implements rate limiting to limit the number of requests made within a certain time period.
type RateLimiter struct { //nolint:govet // To be refactored.
	requestsPerMinute     int
	minimumDelayInSeconds int
	timeOfLastRequest     time.Time
	lock                  sync.Mutex
}

// Options implements the structure of RateLimiter creation options.
type Options struct {
	RequestsPerMinute     int
	MinimumDelayInSeconds int
}

// New creates a new *RateLimiter with the specified *Options.
func New(options *Options) (limiter *RateLimiter) {
	limiter = &RateLimiter{
		requestsPerMinute:     options.RequestsPerMinute,
		minimumDelayInSeconds: options.MinimumDelayInSeconds,
		timeOfLastRequest:     time.Now(),
	}

	return
}

// Wait waits until the next request can be made within the rate limit.
func (limiter *RateLimiter) Wait() {
	limiter.lock.Lock()
	defer limiter.lock.Unlock()

	var timeToSleep time.Duration

	// calculate the time that has elapsed since the last request was made, based on the timeOfLastRequest value of the limiter object.
	timeSinceLastRequest := time.Since(limiter.timeOfLastRequest)

	if limiter.requestsPerMinute > 0 {
		// calculate the minimum interval (in time.Duration units) that should be enforced between requests in order to comply with the rate limiting policy.
		interval := time.Duration(time.Minute.Nanoseconds() / int64(limiter.requestsPerMinute))
		// calculate the amount of time that the program needs to sleep before making another request, in order to comply with the rate limiting policy.
		timeToSleep = interval - timeSinceLastRequest
	}

	minimumDelayInNanoseconds := time.Duration(limiter.minimumDelayInSeconds) * time.Second

	if timeSinceLastRequest < minimumDelayInNanoseconds {
		timeToSleep = minimumDelayInNanoseconds - timeSinceLastRequest
	}

	// sleep
	time.Sleep(timeToSleep)

	// Update the last access time to the current time.
	limiter.timeOfLastRequest = time.Now()
}
