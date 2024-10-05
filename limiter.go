package limiter

import (
	"sync"
	"time"
)

type Limiter struct {
	requestsPerMinute     int
	minimumDelayInSeconds int
	timeOfLastRequest     time.Time
	lock                  sync.Mutex
}

func (limiter *Limiter) Wait() {
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

	if timeToSleep < 0 {
		timeToSleep = 0
	}

	time.Sleep(timeToSleep)

	limiter.timeOfLastRequest = time.Now()
}

type Configuration struct {
	RequestsPerMinute     int
	MinimumDelayInSeconds int
}

func New(cfg *Configuration) (limiter *Limiter) {
	limiter = &Limiter{
		requestsPerMinute:     cfg.RequestsPerMinute,
		minimumDelayInSeconds: cfg.MinimumDelayInSeconds,
		timeOfLastRequest:     time.Now(),
	}

	return
}
