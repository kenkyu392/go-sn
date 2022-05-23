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

// Unix ...
func Unix() Rule {
	return func() []rune {
		return []rune(strconv.Itoa(int(timeNow().Unix())))
	}
}

// UnixMilli ...
func UnixMilli() Rule {
	return func() []rune {
		return []rune(strconv.Itoa(int(timeNow().UnixMilli())))
	}
}

// UnixMicro ...
func UnixMicro() Rule {
	return func() []rune {
		return []rune(strconv.Itoa(int(timeNow().UnixMicro())))
	}
}

// UnixNano ...
func UnixNano() Rule {
	return func() []rune {
		return []rune(strconv.Itoa(int(timeNow().UnixNano())))
	}
}

// Year ...
func Year() Rule {
	return func() []rune {
		return []rune(strconv.Itoa(timeNow().Year()))
	}
}
