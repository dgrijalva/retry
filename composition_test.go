package retry

import (
	"testing"
)

func TestAllComposition(t *testing.T) {
	cases := []int{0, 1, 2, 5, 20, 100}

	for i := 1; i < len(cases); i++ {
		a, b := cases[i-1], cases[i]
		strategry := All{&SimpleStrategy{Tries: a}, &SimpleStrategy{Tries: b}}
		tryCase(t, strategry, testCase{
			name:     []int{a, b},
			attempts: b + 1,
			minimum:  a,
			maximum:  a,
		})
	}

}

func TestAnyComposition(t *testing.T) {
	cases := []int{0, 1, 2, 5, 20, 100}

	for i := 1; i < len(cases); i++ {
		a, b := cases[i-1], cases[i]
		strategry := Any{&SimpleStrategy{Tries: a}, &SimpleStrategy{Tries: b}}
		tryCase(t, strategry, testCase{
			name:     []int{a, b},
			attempts: b + 1,
			minimum:  b,
			maximum:  b,
		})
	}

}
