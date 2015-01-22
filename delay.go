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
		timeSince := TimeFunc().Sub(s.lastTime)
		if timeSince < s.Wait {
			SleepFunc(s.Wait - timeSince)
		}
	}

	s.lastTime = TimeFunc()
	return true
}

func (s *DelayStrategy) HasNext() bool {
	return true
}

// Exponential backoff.  No iteration limit, but it gets
// slower every time.  Resettable.
type ExponentialBackoffStrategy struct {
	InitialDelay time.Duration
	MaxDelay     time.Duration // default: no limit
	count        float64
	lastTime     time.Time
}

func (s *ExponentialBackoffStrategy) Next() bool {
	if !s.lastTime.IsZero() {
		// Calculate the next delay
		nextDelay := time.Duration(math.Pow(s.count, 2)) * s.InitialDelay
		if s.MaxDelay > 0 && nextDelay > s.MaxDelay {
			nextDelay = s.MaxDelay
		}

		timeSince := TimeFunc().Sub(s.lastTime)
		if timeSince < nextDelay {
			SleepFunc(nextDelay - timeSince)
		}
	}

	s.lastTime = TimeFunc()
	s.count++
	return true
}

func (s *ExponentialBackoffStrategy) HasNext() bool {
	return true
}

func (s *ExponentialBackoffStrategy) Reset() {
	var t time.Time
	s.lastTime = t
	s.count = 0
}

// Maximum time strategy.  No iteration limit.  Limit on max time.
// Timer starts automatically on first try. Resettable
type MaximumTimeStrategy struct {
	Duration  time.Duration
	startTime time.Time
}

func (s *MaximumTimeStrategy) Next() bool {
	if s.startTime.IsZero() {
		s.startTime = TimeFunc()
		return true
	}

	return s.HasNext()
}

func (s *MaximumTimeStrategy) elapsed() time.Duration {
	if s.startTime.IsZero() {
		return 0
	}
	return TimeFunc().Sub(s.startTime)
}

func (s *MaximumTimeStrategy) HasNext() bool {
	return s.elapsed() < s.Duration
}

func (s *MaximumTimeStrategy) Reset() {
	var t time.Time
	s.startTime = t
}
