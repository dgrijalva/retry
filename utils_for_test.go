package retry

import (
	"testing"
	"time"
)

type testCase struct {
	name        interface{}
	attempts    int
	minimum     int
	maximum     int
	minDuration time.Duration
	maxDuration time.Duration
}

func tryCase(t *testing.T, strategy RetryStrategy, test testCase) {
	var i int
	start := time.Now()
	for i = 0; i < test.attempts; i++ {
		if d := time.Now().Sub(start); test.maxDuration > 0 && d >= test.maxDuration && strategy.HasNext() {
			t.Errorf("[%v] Unexpected HasNext after time: %v max: %v", test.name, d, test.maxDuration)
		}

		if test.maximum > 0 && i >= test.maximum && strategy.HasNext() {
			t.Errorf("[%v] Unexpected HasNext after attempts: %v maximum: %v", test.name, i, test.maximum)
		}

		next := strategy.Next()

		if d := time.Now().Sub(start); test.maxDuration > 0 && d >= test.maxDuration && next {
			t.Errorf("[%v] Unexpected Next after time: %v max: %v", test.name, d, test.maxDuration)
		}

		if test.maximum > 0 && i > test.maximum && next {
			t.Errorf("[%v] Unexpected Next after attempts: %v maximum: %v", test.name, i, test.maximum)
		}

		if !next {
			break
		}
	}

	if i < test.minimum {
		t.Errorf("[%v] Did not reach minimum attempts: %v minimum: %v", test.name, i, test.minimum)
	}

	if d := time.Now().Sub(start); test.minDuration > 0 && d < test.minDuration {
		t.Errorf("[%v] Test too less time: %v than the minimum", test.name, d, test.minDuration)
	}

}
