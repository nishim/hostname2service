package h2s

import "testing"

func TestUndefined(t *testing.T) {
	e := "undefined"
	a := Detect("qwerty")
	if e != a {
		t.Errorf("expected: %v, actual: %v", e, a)
	}
}
