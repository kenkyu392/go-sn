package rule_test

import (
	"reflect"
	"testing"

	"github.com/kenkyu392/go-sn/rule"
)

func TestRepeatList(t *testing.T) {
	gotA := func() [][]rune {
		out := make([][]rune, 0)
		a := rule.RepeatList([]string{"Dog", "Cat", "Fox", "Rat"})
		for i := 0; i < 6; i++ {
			out = append(out, a())
		}
		return out
	}()
	gotB := func() [][]rune {
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
	if !reflect.DeepEqual(gotA, want) {
		t.Errorf("got: %v, want: %v", gotA, want)
	}
	if !reflect.DeepEqual(gotB, want) {
		t.Errorf("got: %v, want: %v", gotB, want)
	}
}
