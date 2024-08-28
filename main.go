package main

import (
	"fmt"
)

func ParseJSON(jsonString string) (interface{}, error) {
	tokens, err := Lex(jsonString)
	if err != nil {
		return nil, fmt.Errorf("lexing failed: %v", err)
	}

	parser := NewParser(tokens)
	result, err := parser.Parse()
	if err != nil {
		return nil, fmt.Errorf("parsing failed: %v", err)
	}

	return result, nil
}
