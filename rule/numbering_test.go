package rule_test

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"

	"github.com/kenkyu392/go-sn/rule"
)

func TestNumbering(t *testing.T) {
	rule := rule.Numbering()
	for i := 0; i < 11; i++ {
		if got, want := rule(), []rune(strconv.Itoa(i)); !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	}
}

func TestNumberingWithPadding(t *testing.T) {
	rule := rule.NumberingWithPadding(6)
	for i := 0; i < 11; i++ {
		if got, want := rule(), []rune(fmt.Sprintf("%06d", i)); !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	}
}

func TestCounter(t *testing.T) {
	t.Run("case=default", func(t *testing.T) {
		c := &rule.Counter{
			Start:   -1,
			Size:    -1,
			Padding: -1,
		}
		a := c.NumberingWithPadding()
		b := rule.Numbering()
		for i := 0; i < 11; i++ {
			if got, want := a(), b(); !reflect.DeepEqual(got, want) {
				t.Errorf("got: %v, want: %v", got, want)
			}
		}
	})

	t.Run("case=custom", func(t *testing.T) {
		c := &rule.Counter{
			Start:   2,
			Size:    2,
			Padding: 4,
		}
		a := c.NumberingWithPadding()
		for i := 2; i < 14; i += 2 {
			if got, want := a(), []rune(fmt.Sprintf("%04d", i)); !reflect.DeepEqual(got, want) {
				t.Errorf("got: %v, want: %v", got, want)
			}
		}
	})
}
