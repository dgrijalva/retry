package retry

import (
	"testing"
	"time"
)

// Override time for test purposes
var testTime time.Time

func init() {
	TimeFunc = func() time.Time {
		return testTime
	}
	SleepFunc = func(d time.Duration) {
		testTime = testTime.Add(d)
	}
}

type testCase struct {
	name        interface{}
	attempts    int
	minimum     int
	maximum     int
	minDuration time.Duration
	maxDuration time.Duration
	step        time.Duration // how far to step time forward after each iteration
}

func tryCase(t *testing.T, strategy Strategy, test testCase) {

	// Time
	start := time.Now()
	testTime = start
	timeStep := time.Microsecond
	if test.step > 0 {
		timeStep = test.step
	}

	var i int
	for i = 0; i < test.attempts; i++ {
		if d := TimeFunc().Sub(start); test.maxDuration > 0 && d >= test.maxDuration && strategy.HasNext() {
			t.Errorf("[%v] Unexpected HasNext after time: %v max: %v", test.name, d, test.maxDuration)
		}

		if test.maximum > 0 && i >= test.maximum && strategy.HasNext() {
			t.Errorf("[%v] Unexpected HasNext after attempts: %v maximum: %v", test.name, i, test.maximum)
		}

		next := strategy.Next()

		if d := TimeFunc().Sub(start); test.maxDuration > 0 && d >= test.maxDuration && next {
			t.Errorf("[%v] Unexpected Next after time: %v max: %v", test.name, d, test.maxDuration)
		}

		if test.maximum > 0 && i > test.maximum && next {
			t.Errorf("[%v] Unexpected Next after attempts: %v maximum: %v", test.name, i, test.maximum)
		}

		if !next {
			break
		}

		testTime = testTime.Add(timeStep)
	}

	if i < test.minimum {
		t.Errorf("[%v] Did not reach minimum attempts: %v minimum: %v", test.name, i, test.minimum)
	}

	if d := TimeFunc().Sub(start); test.minDuration > 0 && d < test.minDuration {
		t.Errorf("[%v] Test took less time: %v than the minimum: %v", test.name, d, test.minDuration)
	}

}
