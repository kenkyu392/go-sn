package rule

import (
	"strconv"
	"time"
)

var timeNow = time.Now

// SetTimeNow ...
func SetTimeNow(t func() time.Time) {
	timeNow = t
}

// Time ...
func Time(layout string) Rule {
	return func() []rune {
		return []rune(timeNow().Format(layout))
	}
}

// UnixNano ...
func UnixNano() Rule {
	return func() []rune {
		return []rune(strconv.Itoa(int(timeNow().UnixNano())))
	}
}

// UnixMicro ...
func UnixMicro() Rule {
	return func() []rune {
		return []rune(strconv.Itoa(int(timeNow().UnixMicro())))
	}
}

// UnixMilli ...
func UnixMilli() Rule {
	return func() []rune {
		return []rune(strconv.Itoa(int(timeNow().UnixMilli())))
	}
}
