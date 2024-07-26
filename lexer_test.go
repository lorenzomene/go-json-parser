package main

import (
	"fmt"
	"testing"
)

func TestLex(t *testing.T) {
	result, err := Lex("{}")
	if err != nil {
		fmt.Println(err)
	}
	expected := []Token{}

	if !resultEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func resultEqual(a, b []Token) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
