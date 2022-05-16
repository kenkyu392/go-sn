package rule_test

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/kenkyu392/go-sn/rule"
)

func TestRandomAlphabet(t *testing.T) {
	testRule(t, rule.RandomAlphabet(4), [][]rune{
		[]rune("uYIZ"),
		[]rune("abEa"),
		[]rune("ScMg"),
		[]rune("CyWE"),
		[]rune("WhiG"),
		[]rune("MaHu"),
		[]rune("hceE"),
		[]rune("CepM"),
		[]rune("pOTQ"),
		[]rune("plYJ"),
		[]rune("oxfk"),
	})
}

func TestRandomUppercaseAlphabet(t *testing.T) {
	testRule(t, rule.RandomUppercaseAlphabet(4), [][]rune{
		[]rune("CUBY"),
		[]rune("HIZZ"),
		[]rune("KAKB"),
		[]rune("LEEA"),
		[]rune("NSOC"),
		[]rune("PMIG"),
		[]rune("NCKY"),
		[]rune("RWXE"),
		[]rune("VWSH"),
		[]rune("OIPG"),
		[]rune("ZMGA"),
	})
}

func TestRandomLowercaseAlphabet(t *testing.T) {
	testRule(t, rule.RandomLowercaseAlphabet(4), [][]rune{
		[]rune("cuby"),
		[]rune("hizz"),
		[]rune("kakb"),
		[]rune("leea"),
		[]rune("nsoc"),
		[]rune("pmig"),
		[]rune("ncky"),
		[]rune("rwxe"),
		[]rune("vwsh"),
		[]rune("oipg"),
		[]rune("zmga"),
	})
}

func TestRandomNumeric(t *testing.T) {
	testRule(t, rule.RandomNumeric(4), [][]rune{
		[]rune("4436"),
		[]rune("5677"),
		[]rune("8887"),
		[]rune("9826"),
		[]rune("1004"),
		[]rune("5860"),
		[]rune("1620"),
		[]rune("1278"),
		[]rune("5483"),
		[]rune("0094"),
		[]rune("7468"),
	})
}

func TestRandomAlphabetAndNumeric(t *testing.T) {
	testRule(t, rule.RandomAlphabetAndNumeric(4), [][]rune{
		[]rune("uY6Z"),
		[]rune("ab8a"),
		[]rune("0C80"),
		[]rune("6YWe"),
		[]rune("4304"),
		[]rune("MaHu"),
		[]rune("h8eE"),
		[]rune("c4pm"),
		[]rune("PO1q"),
		[]rune("7LYj"),
		[]rune("o1F6"),
	})
}

func TestRandomList(t *testing.T) {
	testRule(t, rule.RandomList([]string{"Dog", "Cat", "Fox", "Rat"}), [][]rune{
		[]rune("Fox"),
		[]rune("Fox"),
		[]rune("Cat"),
		[]rune("Fox"),
		[]rune("Rat"),
		[]rune("Dog"),
		[]rune("Rat"),
		[]rune("Cat"),
		[]rune("Dog"),
		[]rune("Dog"),
		[]rune("Dog"),
	})
}

func testRule(t *testing.T, rule rule.Rule, wants [][]rune) {
	rand.Seed(0)
	for _, want := range wants {
		if got := rule(); !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", string(got), string(want))
		}
	}
}
