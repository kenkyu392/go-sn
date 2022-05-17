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
		got := rule()
		want := []rune(strconv.Itoa(i))
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	}
}

func TestNumberingWithPadding(t *testing.T) {
	rule := rule.NumberingWithPadding(6)
	for i := 0; i < 11; i++ {
		got := rule()
		want := []rune(fmt.Sprintf("%06d", i))
		if !reflect.DeepEqual(got, want) {
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
		testRule(t, c.NumberingWithPadding(), [][]rune{
			[]rune("0002"),
			[]rune("0004"),
			[]rune("0006"),
			[]rune("0008"),
			[]rune("0010"),
			[]rune("0012"),
		})
	})
}
