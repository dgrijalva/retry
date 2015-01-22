package retry

import (
	"testing"
	"time"
)

func TestDelayStrategy(t *testing.T) {
	for _, test := range []time.Duration{0, time.Microsecond, 10 * time.Millisecond} {
		tryCase(t, &DelayStrategy{Wait: test}, testCase{
			name:        test,
			attempts:    3,
			minDuration: 2 * test,
			maxDuration: 3 * test,
		})
	}
}

func TestExponentialBackoffStrategy(t *testing.T) {
	for _, test := range []struct {
		iterations int
		initial    time.Duration
		maxWait    time.Duration
		duration   time.Duration
	}{
		{2, time.Microsecond, 0, time.Microsecond},
		{3, time.Millisecond, 0, 5 * time.Millisecond},
		{5, time.Millisecond, 0, 30 * time.Millisecond},
		{5, time.Millisecond, 2 * time.Millisecond, 8 * time.Millisecond},
	} {
		tryCase(t, &ExponentialBackoffStrategy{InitialDelay: test.initial, MaxDelay: test.maxWait}, testCase{
			name:        test,
			attempts:    test.iterations,
			minDuration: test.duration,
			maxDuration: test.duration + time.Microsecond,
			step:        test.initial,
		})
	}
}

func TestMaximumTimeStrategy(t *testing.T) {
	for _, test := range []struct {
		attempts   int
		iterations int
		duration   time.Duration
	}{
		{2, 1, time.Millisecond},
		{10, 5, 5 * time.Millisecond},
	} {
		tryCase(t, &MaximumTimeStrategy{Duration: test.duration}, testCase{
			name:        test,
			attempts:    test.attempts,
			minimum:     test.iterations,
			maximum:     test.iterations,
			minDuration: test.duration,
			maxDuration: test.duration,
			step:        time.Millisecond,
		})
	}
}
