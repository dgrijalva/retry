package retry

type SimpleStrategy struct {
	Tries int
	count int
}

// Type validation
var _ RetryStrategy = &SimpleStrategy{}

func (s *SimpleStrategy) Next() bool {
	if s.count < s.Tries {
		s.count++
		return true
	}
	return false
}

func (s *SimpleStrategy) HasNext() bool {
	return s.count < s.Tries
}
