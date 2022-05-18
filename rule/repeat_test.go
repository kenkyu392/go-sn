package rule_test

import (
	"reflect"
	"testing"

	"github.com/kenkyu392/go-sn/rule"
)

func TestRepeat(t *testing.T) {
	got := func() [][]rune {
		out := make([][]rune, 0)
		a := rule.Repeat([]string{"Dog", "Cat", "Fox", "Rat"})
		for i := 0; i < 6; i++ {
			out = append(out, a())
		}
		return out
	}()
	want := [][]rune{
		[]rune("Dog"),
		[]rune("Cat"),
		[]rune("Fox"),
		[]rune("Rat"),
		[]rune("Dog"),
		[]rune("Cat"),
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
