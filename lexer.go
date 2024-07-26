package main

import (
	"errors"
	"fmt"
	// "unicode"
)

func Lex(s string) ([]Token, error) {
	var tokens []Token
	for runes := []rune(s); len(runes) > 0; {
		char := runes[0]

		switch char {
		case '{':
			tokens = append(tokens, Token{Type: LEFT_BRACE})
			runes = runes[1:]
		case '}':
			tokens = append(tokens, Token{Type: RIGHT_BRACE})
			runes = runes[1:]
		case ':':
			tokens = append(tokens, Token{Type: COLON})
			runes = runes[1:]
		case ',':
			tokens = append(tokens, Token{Type: COMMA})
			runes = runes[1:]
		case '"': //string start
			j := 1
			//peeking ahead to get the values of the string
			for j < len(runes) && runes[j] != '"' {
				j++
			}
			if j >= len(runes) { //string end not found
				return nil, errors.New("unterminated string")
			}
			tokens = append(tokens, Token{Type: COMMA})
			runes = runes[1:]
		}
		runes = runes[1:]
	}
	fmt.Println(tokens)
	return tokens, nil
}
