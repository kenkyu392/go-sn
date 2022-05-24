package rule

import (
	"sync/atomic"
)

// RepeatList creates a rule that repeatedly returns the items of a list in order.
func RepeatList(list []string) Rule {
	var n int64 = 0
	return func() []rune {
		index := atomic.LoadInt64(&n)
		if index+1 >= int64(len(list)) {
			atomic.StoreInt64(&n, 0)
		} else {
			atomic.AddInt64(&n, 1)
		}
		return []rune(list[index])
	}
}

// Repeat creates a rule that repeatedly returns the items of a list in order.
// Deprecated: Repeat exists for historical compatibility and should not be used.
// RepeatList should be used instead.
func Repeat(list []string) Rule {
	return RepeatList(list)
}
