package retry

import (
	"testing"
	"time"
)

func TestDelayStrategy(t *testing.T) {
	for _, test := range []time.Duration{0, 1, time.Microsecond, 10 * time.Millisecond} {
		tryCase(t, &DelayStrategy{Wait: test}, testCase{
			name:        test,
			attempts:    3,
			minDuration: 2 * test,
			maxDuration: (3 * test) + (time.Millisecond / 2),
		})
	}
}

func TestExponentialBackoffStrategy(t *testing.T) {
	for _, test := range []struct {
		iterations int
		initial    time.Duration
		duration   time.Duration
	}{
		{1, 1, 1},
		{2, time.Microsecond, time.Microsecond},
		{3, time.Millisecond, 5 * time.Millisecond},
		{5, time.Millisecond, 29 * time.Millisecond},
	} {
		tryCase(t, &ExponentialBackoffStrategy{InitialDelay: test.initial}, testCase{
			name:        test,
			attempts:    test.iterations,
			minDuration: test.duration,
			maxDuration: time.Duration(float64(test.duration)*1.4) + (time.Millisecond / 4),
		})
	}
}
