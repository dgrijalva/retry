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
