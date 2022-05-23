package rule_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/kenkyu392/go-sn/rule"
)

func TestTime(t *testing.T) {
	rule.SetTimeNow(func() time.Time {
		return time.Date(2021, 1, 1, 1, 1, 1, 1, time.UTC)
	})

	got := rule.Time(time.RFC3339Nano)()
	want := []rune("2021-01-01T01:01:01.000000001Z")
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestUnixXXXX(t *testing.T) {
	rule.SetTimeNow(func() time.Time {
		return time.Date(2021, 1, 1, 1, 1, 1, 1, time.UTC)
	})

	t.Run("case=UnixNano", func(t *testing.T) {
		got := rule.UnixNano()()
		want := []rune("1609462861000000001")
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("case=UnixMicro", func(t *testing.T) {
		got := rule.UnixMicro()()
		want := []rune("1609462861000000")
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("case=UnixMilli", func(t *testing.T) {
		got := rule.UnixMilli()()
		want := []rune("1609462861000")
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("case=Unix", func(t *testing.T) {
		got := rule.Unix()()
		want := []rune("1609462861")
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

func TestYear(t *testing.T) {
	rule.SetTimeNow(func() time.Time {
		return time.Date(2021, 1, 1, 1, 1, 1, 1, time.UTC)
	})

	got := rule.Year()()
	want := []rune("2021")
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
