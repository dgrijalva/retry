package retry

// A composite strategy.  In order for Next or HasNext to
// succeed, all of the included strategies must succeed
type All []Strategy

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

// A composite strategy.  In order for Next or HasNext to
// succeed, any one of the included strategies must succeed
type Any []Strategy

func (s Any) Next() bool {
	// Call all strategies even one returns true
	// otherwise, they might lose count
	var succ = false
	for _, ss := range s {
		if ss.Next() {
			succ = true
		}
	}
	return succ
}

func (s Any) HasNext() bool {
	for _, ss := range s {
		if ss.HasNext() {
			return true
		}
	}
	return false
}
