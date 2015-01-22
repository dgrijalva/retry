package retry

import (
	"testing"
	"time"
)

type testCase struct {
	name        interface{}
	attempts    int
	limit       int
	minDuration time.Duration
	maxDuration time.Duration
}

func tryCase(t *testing.T, strategy RetryStrategy, test testCase) {
	start := time.Now()
	for i := 0; i < test.attempts; i++ {
		if d := time.Now().Sub(start); test.maxDuration > 0 && d >= test.maxDuration && strategy.HasNext() {
			t.Errorf("[%v] Unexpected HasNext after time: %v max: %v", test.name, d, test.maxDuration)
		}

		if test.limit > 0 && i >= test.limit && strategy.HasNext() {
			t.Errorf("[%v] Unexpected HasNext after attempts: %v limit: %v", test.name, i, test.limit)
		}

		next := strategy.Next()

		if d := time.Now().Sub(start); test.maxDuration > 0 && d >= test.maxDuration && next {
			t.Errorf("[%v] Unexpected Next after time: %v max: %v", test.name, d, test.maxDuration)
		}

		if test.limit > 0 && i > test.limit && next {
			t.Errorf("[%v] Unexpected Next after attempts: %v limit: %v", test.name, i, test.limit)
		}
	}

	if d := time.Now().Sub(start); test.minDuration > 0 && d < test.minDuration {
		t.Errorf("[%v] Test too less time: %v than the minimum", test.name, d, test.minDuration)
	}

}
