package retry

type RetryStrategy interface {
	Next() bool
	HasNext() bool
}

type ResettableRetryStrategy interface {
	RetryStrategy
	Reset()
}

func Do(strategy RetryStrategy, action func() bool) bool {
	for strategy.Next() {
		if succ := action(); succ {
			return succ
		}
	}
	return false
}
