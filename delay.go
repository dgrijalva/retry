package retry

import (
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
