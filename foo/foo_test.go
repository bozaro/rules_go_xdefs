package foo_test

import (
	"testing"

	"github.com/bozaro/rules_go_xdefs/bar"
)

func TestBar(t *testing.T) {
	Equal(t, "Bar", bar.Bar)
}

func TestBaz(t *testing.T) {
	Equal(t, "Baz", bar.Baz)
}

func TestQux(t *testing.T) {
	Equal(t, "Qux", bar.Qux)
}

func Equal(t *testing.T, expected string, actual string) bool {
	if expected != actual {
		t.Errorf("Not equal: \n"+
			"expected: %s\n"+
			"actual  : %s", expected, actual)
	}
	return true
}
