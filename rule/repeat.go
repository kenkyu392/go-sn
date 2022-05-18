package rule

import "sync/atomic"

// Repeat creates a rule that repeatedly returns the items of a list in order.
func Repeat(list []string) Rule {
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
