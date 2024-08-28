package main

import (
	"reflect"
	"testing"
)

func TestParseObject(t *testing.T) {
	//{"name": "John", "age": 30, "isStudent": false}
	tokens := []Token{
		{Type: LEFT_BRACE, Value: "{"},
		{Type: STRING, Value: "name"},
		{Type: COLON, Value: ":"},
		{Type: STRING, Value: "John"},
		{Type: COMMA, Value: ","},
		{Type: STRING, Value: "age"},
		{Type: COLON, Value: ":"},
		{Type: NUMBER, Value: "30"},
		{Type: COMMA, Value: ","},
		{Type: STRING, Value: "isStudent"},
		{Type: COLON, Value: ":"},
		{Type: FALSE, Value: "false"},
		{Type: RIGHT_BRACE, Value: "}"},
	}

	expected := map[string]interface{}{
		"name":      "John",
		"age":       float64(30),
		"isStudent": false,
	}

	parser := NewParser(tokens)
	result, err := parser.Parse()
	if err != nil {
		t.Fatalf("Parse() returned an error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Parse() = %v, expected %v", result, expected)
	}
}

func TestParseArray(t *testing.T) {
	//[1, "two", false, null]
	tokens := []Token{
		{Type: LEFT_BRACKET, Value: "["},
		{Type: NUMBER, Value: "1"},
		{Type: COMMA, Value: ","},
		{Type: STRING, Value: "two"},
		{Type: COMMA, Value: ","},
		{Type: FALSE, Value: "false"},
		{Type: COMMA, Value: ","},
		{Type: NULL, Value: "null"},
		{Type: RIGHT_BRACKET, Value: "]"},
	}

	expected := []interface{}{
		float64(1),
		"two",
		false,
		nil,
	}

	parser := NewParser(tokens)
	result, err := parser.Parse()
	if err != nil {
		t.Fatalf("Parse() returned an error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Parse() = %v, expected %v", result, expected)
	}
}

func TestParseNestedObject(t *testing.T) {
	//{"person": {"name": "John", "age": 30}}
	tokens := []Token{
		{Type: LEFT_BRACE, Value: "{"},
		{Type: STRING, Value: "person"},
		{Type: COLON, Value: ":"},
		{Type: LEFT_BRACE, Value: "{"},
		{Type: STRING, Value: "name"},
		{Type: COLON, Value: ":"},
		{Type: STRING, Value: "John"},
		{Type: COMMA, Value: ","},
		{Type: STRING, Value: "age"},
		{Type: COLON, Value: ":"},
		{Type: NUMBER, Value: "30"},
		{Type: RIGHT_BRACE, Value: "}"},
		{Type: RIGHT_BRACE, Value: "}"},
	}

	expected := map[string]interface{}{
		"person": map[string]interface{}{
			"name": "John",
			"age":  float64(30),
		},
	}

	parser := NewParser(tokens)
	result, err := parser.Parse()
	if err != nil {
		t.Fatalf("Parse() returned an error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Parse() = %v, expected %v", result, expected)
	}
}

func TestParseNestedArray(t *testing.T) {
	//{"numbers": [1, 2, [3, 4]]}
	tokens := []Token{
		{Type: LEFT_BRACE, Value: "{"},
		{Type: STRING, Value: "numbers"},
		{Type: COLON, Value: ":"},
		{Type: LEFT_BRACKET, Value: "["},
		{Type: NUMBER, Value: "1"},
		{Type: COMMA, Value: ","},
		{Type: NUMBER, Value: "2"},
		{Type: COMMA, Value: ","},
		{Type: LEFT_BRACKET, Value: "["},
		{Type: NUMBER, Value: "3"},
		{Type: COMMA, Value: ","},
		{Type: NUMBER, Value: "4"},
		{Type: RIGHT_BRACKET, Value: "]"},
		{Type: RIGHT_BRACKET, Value: "]"},
		{Type: RIGHT_BRACE, Value: "}"},
	}

	expected := map[string]interface{}{
		"numbers": []interface{}{
			float64(1),
			float64(2),
			[]interface{}{
				float64(3),
				float64(4),
			},
		},
	}

	parser := NewParser(tokens)
	result, err := parser.Parse()
	if err != nil {
		t.Fatalf("Parse() returned an error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Parse() = %v, expected %v", result, expected)
	}
}

func TestParseEmptyObject(t *testing.T) {
	//{}
	tokens := []Token{
		{Type: LEFT_BRACE, Value: "{"},
		{Type: RIGHT_BRACE, Value: "}"},
	}

	expected := map[string]interface{}{}

	parser := NewParser(tokens)
	result, err := parser.Parse()
	if err != nil {
		t.Fatalf("Parse() returned an error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Parse() = %v, expected %v", result, expected)
	}
}

func TestParseEmptyArray(t *testing.T) {
	// Example JSON: []
	tokens := []Token{
		{Type: LEFT_BRACKET, Value: "["},
		{Type: RIGHT_BRACKET, Value: "]"},
	}

	parser := NewParser(tokens)
	result, err := parser.Parse()
	if err != nil {
		t.Fatalf("Parse() returned an error: %v", err)
	}

	// Check if the result is an empty slice
	if arr, ok := result.([]interface{}); !ok || len(arr) != 0 {
		t.Errorf("Parse() = %v, expected an empty array", result)
	}
}

func TestParseInvalidToken(t *testing.T) {
	//{"key": invalid}
	tokens := []Token{
		{Type: LEFT_BRACE, Value: "{"},
		{Type: STRING, Value: "key"},
		{Type: COLON, Value: ":"},
		{Type: NUMBER, Value: "123abc"},
		{Type: RIGHT_BRACE, Value: "}"},
	}

	parser := NewParser(tokens)
	_, err := parser.Parse()
	if err == nil {
		t.Fatalf("Parse() expected to return an error for invalid token, but got nil")
	}
}
