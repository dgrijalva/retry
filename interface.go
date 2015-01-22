package retry

type RetryStrategy interface {
	Next() bool
	HasNext() bool
}

func Do(strategy RetryStrategy, action func() bool) bool {
	for strategy.Next() {
		if succ := action(); succ {
			return succ
		}
	}
	return false
}
