package main

import "testing"

func TestAdd(t *testing.T) {
	var expected int = 12

	var actual int = Add(6, 6)

	if expected != actual {
		t.Errorf("expected: %d got: %d", expected, actual)
	}
}
