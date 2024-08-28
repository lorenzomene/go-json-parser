package main

import (
	"fmt"
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
			if !deepEqual(result, tt.expected) {
				fmt.Printf("Result: %#v\n", result)
				fmt.Printf("Expected: %#v\n", tt.expected)
				t.Errorf("ParseJSON() = %#v, expected %#v", result, tt.expected)
			}
		})
	}
}

func deepEqual(a, b interface{}) bool {
	if reflect.DeepEqual(a, b) {
		return true
	}

	switch av := a.(type) {
	case []interface{}:
		bv, ok := b.([]interface{})
		if !ok {
			return false
		}

		if len(av) == 0 && len(bv) == 0 {
			return true
		}
		return reflect.DeepEqual(av, bv)
	case map[string]interface{}:
		bv, ok := b.(map[string]interface{})
		if !ok {
			return false
		}

		if len(av) == 0 && len(bv) == 0 {
			return true
		}
		return reflect.DeepEqual(av, bv)
	}

	return false
}
