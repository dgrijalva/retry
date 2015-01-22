package retry

// Retry forever.  No waiting.
type AlwaysRetryStrategy struct{}

func (s *AlwaysRetryStrategy) Next() bool {
	return true
}

func (s *AlwaysRetryStrategy) HasNext() bool {
	return true
}

// Try up to a fixed number of times
type CountStrategy struct {
	Tries int
	count int
}

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

func (s *CountStrategy) Reset() {
	s.count = 0
}
