package retry

type AlwaysRetryStrategy struct{}

func (s *AlwaysRetryStrategy) Next() bool {
	return true
}

func (s *AlwaysRetryStrategy) HasNext() bool {
	return true
}

type CountStrategy struct {
	Tries int
	count int
}

// Type validation
var _ Strategy = &CountStrategy{}

func (s *CountStrategy) Next() bool {
	if s.count < s.Tries {
		s.count++
		return true
	}
	return false
}

func (s *CountStrategy) HasNext() bool {
	return s.count < s.Tries
}
