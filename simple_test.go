package retry

import (
	"testing"
)

func TestCountStrategy(t *testing.T) {
	for _, test := range []int{0, 1, 2, 5, 100} {
		tryCase(t, &CountStrategy{Tries: test}, testCase{
			name:     test,
			attempts: test + 10,
			minimum:  test,
			maximum:  test,
		})
	}
}
