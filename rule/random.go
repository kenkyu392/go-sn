package rule

import (
	"math/rand"
	"strconv"
)

// RandomAlphabet creates a rule that returns a random alphabet (a-zA-Z)
// with the number of characters specified by length.
func RandomAlphabet(length int) Rule {
	return func() []rune {
		var out = make([]rune, length)
		for i := 0; i < length; i++ {
			switch rand.Intn(2) {
			case 0:
				out[i] = 'a' + rune(rand.Intn(26))
			case 1:
				out[i] = 'A' + rune(rand.Intn(26))
			}
		}
		return out
	}
}

// RandomUppercaseAlphabet creates a rule that returns a random alphabet (A-Z)
// with the number of characters specified by length.
func RandomUppercaseAlphabet(length int) Rule {
	return func() []rune {
		var out = make([]rune, length)
		for i := 0; i < length; i++ {
			out[i] = 'A' + rune(rand.Intn(26))
		}
		return out
	}
}

// RandomLowercaseAlphabet creates a rule that returns a random alphabet (a-z)
// with the number of characters specified by length.
func RandomLowercaseAlphabet(length int) Rule {
	return func() []rune {
		var out = make([]rune, length)
		for i := 0; i < length; i++ {
			out[i] = 'a' + rune(rand.Intn(26))
		}
		return out
	}
}

// RandomNumeric creates a rule that returns a random numeric (0-9)
// with the number of characters specified by length.
func RandomNumeric(length int) Rule {
	return func() []rune {
		var out = make([]rune, length)
		for i := 0; i < length; i++ {
			out[i] = []rune(strconv.Itoa(rand.Intn(10)))[0]
		}
		return out
	}
}

// RandomAlphabetAndNumeric creates a rule that returns a random alphabet and
// numeric (a-zA-Z0-9) with the number of characters specified by length.
func RandomAlphabetAndNumeric(length int) Rule {
	return func() []rune {
		var out = make([]rune, length)
		for i := 0; i < length; i++ {
			switch rand.Intn(3) {
			case 0:
				out[i] = 'a' + rune(rand.Intn(26))
			case 1:
				out[i] = 'A' + rune(rand.Intn(26))
			case 2:
				out[i] = []rune(strconv.Itoa(rand.Intn(10)))[0]
			}
		}
		return out
	}
}

// RandomList creates a rule that returns a list value at random.
func RandomList(list []string) Rule {
	var length = len(list)
	return func() []rune {
		return []rune(list[rand.Intn(length)])
	}
}
