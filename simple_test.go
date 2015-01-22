package retry

import (
	"testing"
)

func TestSimpleStrategy(t *testing.T) {
	for _, test := range []int{0, 1, 2, 5, 100} {
		var count int = 0
		var strategy = &SimpleStrategy{Tries: test}
		Do(strategy, func() bool {
			count++
			if count >= test && strategy.HasNext() {
				t.Errorf("Unexpected HasNext after attempts: %v Limit: %v", count, test)
			}
			return false
		})
		if count != test {
			t.Errorf("Expected attempts: %v Actual: %v", test, count)
		}
	}
}
