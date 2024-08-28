package main

import (
	"reflect"
	"testing"
)

func TestParseJSON(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			input: `{"name": "John Doe", "age": 30, "isStudent": false, "courses": ["Math", "Science"]}`,
			expected: map[string]interface{}{
				"name":      "John Doe",
				"age":       float64(30),
				"isStudent": false,
				"courses": []interface{}{
					"Math",
					"Science",
				},
			},
		},
		{
			input:    `[]`,
			expected: []interface{}{},
		},
		{
			input: `[1, "two", true, null]`,
			expected: []interface{}{
				float64(1),
				"two",
				true,
				nil,
			},
		},
		{
			input: `{"emptyObject": {}, "emptyArray": []}`,
			expected: map[string]interface{}{
				"emptyObject": map[string]interface{}{},
				"emptyArray":  []interface{}{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := ParseJSON(tt.input)
			if err != nil {
				t.Fatalf("ParseJSON() returned an error: %v", err)
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ParseJSON() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestParseJSONInvalid(t *testing.T) {
	invalidJSON := `{"key": invalid}`
	_, err := ParseJSON(invalidJSON)
	if err == nil {
		t.Fatalf("ParseJSON() expected to return an error for invalid JSON, but got nil")
	}
}
