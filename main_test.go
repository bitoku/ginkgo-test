package main

import "testing"

func TestAdd(t *testing.T) {
	if add(1, 1) == 2 { // bad test
		t.Error("Expected 1 + 1 to equal 2")
	}
}
