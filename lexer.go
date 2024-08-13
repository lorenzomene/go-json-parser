package main

import (
	"errors"
	"fmt"
	"unicode"
)

func Lex(s string) ([]Token, error) {
	var tokens []Token
	for runes := []rune(s); len(runes) > 0; {
		char := runes[0]

		switch char {
		case '{':
			tokens = append(tokens, Token{Type: LEFT_BRACE, Value: "{"})
			runes = runes[1:]
		case '}':
			tokens = append(tokens, Token{Type: RIGHT_BRACE, Value: "}"})
			runes = runes[1:]
		case ':':
			tokens = append(tokens, Token{Type: COLON, Value: ":"})
			runes = runes[1:]
		case ',':
			tokens = append(tokens, Token{Type: COMMA, Value: ","})
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
			tokens = append(tokens, Token{Type: STRING, Value: string(runes[1:j])})
			runes = runes[j+1:]
		default:
			// check for numbers and - (negative numbers)
			if unicode.IsDigit(char) || char == '-' {
				j := 0
				if char == '-' {
					j = 1 // skip - sign for digit check (it will still be included to the final token)
				}
				for j < len(runes) && (unicode.IsDigit(runes[j]) || runes[j] == '.') {
					j++
				}
				tokens = append(tokens, Token{Type: NUMBER, Value: string(runes[:j])})
				runes = runes[j:]
			} else if unicode.IsLetter(char) {
				j := 0
				for j < len(runes) && unicode.IsLetter(runes[j]) {
					j++
				}
				word := string(runes[:j])
				switch word {
				case "true":
					tokens = append(tokens, Token{Type: TRUE, Value: word})
				case "false":
					tokens = append(tokens, Token{Type: FALSE, Value: word})
				case "null":
					tokens = append(tokens, Token{Type: NULL, Value: word})
				default:
					return nil, fmt.Errorf("unexpected keyword: %s", word)
				}
				runes = runes[j:]
			} else {
				return nil, fmt.Errorf("unexpected character: %c", char)
			}
		}
		if len(runes) > 0 && (unicode.IsSpace(char) || !unicode.IsLetter(char) && !unicode.IsDigit(char) && char != '"' && char != '{' && char != '}' && char != ':' && char != ',') {
			runes = runes[1:]
		}
	}
	fmt.Println(tokens)
	return tokens, nil
}
