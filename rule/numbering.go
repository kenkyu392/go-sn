package rule

import (
	"fmt"
	"strconv"
	"sync/atomic"
)

// Counter provides methods and options for numbering.
type Counter struct {
	n     int64
	Start int64
	// Size is the value to be incremented.
	// If 0 or less is specified, the value will be 1.
	Size    int64
	Padding int
}

// NumberingWithPadding creates a rule that returns a counted-up number each time it is executed.
// The number returned by this rule will be zero-padded by the number of digits specified in Counter.Padding.
func (c *Counter) NumberingWithPadding() Rule {
	if c.Start < 0 {
		c.Start = 0
	}
	c.n = c.Start
	if c.Size <= 0 {
		c.Size = 1
	}
	if c.Padding < 0 {
		c.Padding = 1
	}
	var format = "%0" + strconv.Itoa(c.Padding) + "d"
	return func() []rune {
		out := []rune(fmt.Sprintf(format, c.n))
		atomic.AddInt64(&c.n, c.Size)
		return out
	}
}

// Numbering creates a rule that returns a counted-up number each time it is executed.
func Numbering() Rule {
	return (&Counter{Start: 0, Size: 1, Padding: 1}).NumberingWithPadding()
}

// NumberingWithPadding creates a rule that returns a counted-up number each time it is executed.
// The number returned by this rule will be zero-padded by the number of digits specified in pad.
func NumberingWithPadding(pad int) Rule {
	return (&Counter{Start: 0, Size: 1, Padding: pad}).NumberingWithPadding()
}
