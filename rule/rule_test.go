package rule_test

import (
	"reflect"
	"testing"

	"github.com/kenkyu392/go-sn/rule"
)

func TestString(t *testing.T) {
	got := rule.String("test")()
	want := []rune("test")
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestMixin(t *testing.T) {
	got := rule.Mixin(
		rule.String("hello"),
		rule.String("world"),
	)()
	want := []rune("helloworld")
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
