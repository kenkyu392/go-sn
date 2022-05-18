package sn_test

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/kenkyu392/go-sn"
	"github.com/kenkyu392/go-sn/rule"
)

func TestGenerator(t *testing.T) {
	rand.Seed(0)

	got := func() [][]rune {
		g := sn.NewGenerator(
			rule.Mixin(
				rule.RandomAlphabet(4),
				rule.String("."),
				rule.RandomAlphabetAndNumeric(6),
			),
			rule.RandomList(
				[]string{
					"example.com",
					"example.net",
					"example.org",
				},
			),
		)
		g.Delimiter = "@"
		out := make([][]rune, 0)
		for i := 0; i < 5; i++ {
			out = append(out, []rune(g.String()))
		}
		return out
	}()

	want := [][]rune{
		[]rune("uYIZ.ab8a0C@example.org"),
		[]rune("inkr.X5SOPz@example.com"),
		[]rune("aHuh.8eEc4p@example.com"),
		[]rune("gHrZ.us9Fus@example.net"),
		[]rune("fkue.8AobH1@example.com"),
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
