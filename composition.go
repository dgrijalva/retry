package retry

type All []RetryStrategy

func (s All) Next() bool {
	for _, ss := range s {
		if !ss.Next() {
			return false
		}
	}
	return true
}

func (s All) HasNext() bool {
	for _, ss := range s {
		if !ss.HasNext() {
			return false
		}
	}
	return true
}

type Any []RetryStrategy

func (s Any) Next() bool {
	for _, ss := range s {
		if ss.Next() {
			return true
		}
	}
	return true
}

func (s All) HasNext() bool {
	for _, ss := range s {
		if ss.HasNext() {
			return true
		}
	}
	return false
}
