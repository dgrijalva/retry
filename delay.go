package retry

import (
	"math"
	"time"
)

// Delay strategy has no limit.  It implements a fixed
// wait time between retries.
type DelayStrategy struct {
	Wait     time.Duration
	lastTime time.Time
}

func (s *DelayStrategy) Next() bool {
	if !s.lastTime.IsZero() {
		timeSince := time.Now().Sub(s.lastTime)
		if timeSince < s.Wait {
			time.Sleep(s.Wait - timeSince)
		}
	}

	s.lastTime = time.Now()
	return true
}

func (s *DelayStrategy) HasNext() bool {
	return true
}

// Exponential backoff.  No iteration limit, but it gets
// slower every time.  Reset clears the count, but does not
// reset the last time.
type ExponentialBackoffStrategy struct {
	InitialDelay time.Duration
	count        float64
	lastTime     time.Time
}

func (s *ExponentialBackoffStrategy) Next() bool {
	if !s.lastTime.IsZero() {
		nextDelay := time.Duration(math.Pow(s.count, 2)) * s.InitialDelay
		timeSince := time.Now().Sub(s.lastTime)
		if timeSince < nextDelay {
			time.Sleep(nextDelay - timeSince)
		}
	}

	s.lastTime = time.Now()
	s.count++
	return true
}

func (s *ExponentialBackoffStrategy) HasNext() bool {
	return true
}

func (s *ExponentialBackoffStrategy) Reset() {
	s.count = 0
}
