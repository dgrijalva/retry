package retry

import (
	"testing"
	"time"
)

func TestIterationsAndTime(t *testing.T) {
	for _, test := range []struct {
		attempts   int
		iterations int
		duration   time.Duration
	}{
		{15, 1, time.Millisecond},
		{10, 30, time.Millisecond},
	} {
		tryCase(t, &All{
			&CountStrategy{Tries: test.iterations},
			&MaximumTimeStrategy{Duration: test.duration},
		}, testCase{
			name:        test,
			attempts:    test.attempts,
			maximum:     test.iterations,
			maxDuration: test.duration,
			step:        time.Millisecond / 10,
		})
	}

}

func TestMinIterationsAndMaxTime(t *testing.T) {
	for _, test := range []struct {
		attempts      int
		minIterations int
		maxIterations int
		duration      time.Duration
		step          time.Duration
	}{
		{1000, 2, 4, time.Millisecond, time.Millisecond / 3},
		{1000, 10, 30, time.Millisecond, time.Millisecond / 9},
	} {
		// At least minIterations
		// At most duration
		// At least step between iterations
		tryCase(t, &Any{
			&CountStrategy{Tries: test.minIterations},
			&All{
				&MaximumTimeStrategy{Duration: test.duration},
				&DelayStrategy{Wait: test.step},
			},
		}, testCase{
			name:        test,
			attempts:    test.attempts,
			minimum:     test.minIterations,
			maximum:     test.maxIterations,
			maxDuration: test.duration,
			step:        time.Millisecond / 100,
		})
	}

}
