package main

import (
	"reflect"
	"testing"
)

func TestLex(t *testing.T) {
	tests := []struct {
		input    string
		expected []Token
	}{
		{
			input: `{"name": "John Doe", "age": 30, "isStudent": false, "email": null}`,
			expected: []Token{
				{Type: LEFT_BRACE, Value: "{"},
				{Type: STRING, Value: "name"},
				{Type: COLON, Value: ":"},
				{Type: STRING, Value: "John Doe"},
				{Type: COMMA, Value: ","},
				{Type: STRING, Value: "age"},
				{Type: COLON, Value: ":"},
				{Type: NUMBER, Value: "30"},
				{Type: COMMA, Value: ","},
				{Type: STRING, Value: "isStudent"},
				{Type: COLON, Value: ":"},
				{Type: FALSE, Value: "false"},
				{Type: COMMA, Value: ","},
				{Type: STRING, Value: "email"},
				{Type: COLON, Value: ":"},
				{Type: NULL, Value: "null"},
				{Type: RIGHT_BRACE, Value: "}"},
			},
		},
		{
			input: `{"empty": null}`,
			expected: []Token{
				{Type: LEFT_BRACE, Value: "{"},
				{Type: STRING, Value: "empty"},
				{Type: COLON, Value: ":"},
				{Type: NULL, Value: "null"},
				{Type: RIGHT_BRACE, Value: "}"},
			},
		},
		{
			input: `{"flag": true}`,
			expected: []Token{
				{Type: LEFT_BRACE, Value: "{"},
				{Type: STRING, Value: "flag"},
				{Type: COLON, Value: ":"},
				{Type: TRUE, Value: "true"},
				{Type: RIGHT_BRACE, Value: "}"},
			},
		},
		{
			input: `{"pi": 3.14}`,
			expected: []Token{
				{Type: LEFT_BRACE, Value: "{"},
				{Type: STRING, Value: "pi"},
				{Type: COLON, Value: ":"},
				{Type: NUMBER, Value: "3.14"},
				{Type: RIGHT_BRACE, Value: "}"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			tokens, err := Lex(tt.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(tokens, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, tokens)
			}
		})
	}
}
