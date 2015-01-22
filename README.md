An interface for and collection of various implementations of retry behavior, designed to be easy to use and composable.

The interface is inspired by [`AttemptStrategy`](http://godoc.org/github.com/crowdmob/goamz/aws#AttemptStrategy) in `github.com/crowdmob/goamz/aws`.  On top of being fairly useful, this library is a good illustration of the usefulness of interfaces in go.  Use a single strategy or combine a few together to easily accomplish what is normally a bit of a chore.

Some examples:

```go
// A very basic example.
strategy := &retry.SimpleStrategy{Tries: 3}
for strategy.Next() {
	trySomething()
}
```

```go
// Compose a few strategies together
// Try not more than 3 times, with a 100ms delay between attempts
strategy := &retry.All{
	&retry.SimpleStrategy{Tries: 3},
	&retry.DelayStrategy{Delay: 100 * time.Millisecond},
}

for strategy.Next() {
	trySomething()
}
```

```go
// Separate retry logic from a more complex function
func doComplexThings(strategy retry.Strategy)bool{
	for strategy.Next() {
		if success := trySomthingThatMightFail(); success {
			return true
		}
	}
	return false
}

doComplexThings(&retry.SimpleStrategy{Tries: 3})
```

