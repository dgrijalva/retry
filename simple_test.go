package retry

import (
	"testing"
)

func TestSimpleStrategy(t *testing.T) {
	for _, test := range []int{0, 1, 2, 5, 100} {
		tryCase(t, &SimpleStrategy{Tries: test}, testCase{
			name:     test,
			attempts: test + 1,
			limit:    test,
		})
	}
}
