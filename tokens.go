package main

type TokenType int

const (
	LEFT_BRACE TokenType = iota
	RIGHT_BRACE
	COLON
	COMMA
	STRING
	NUMBER
)

type Token struct {
	Type  TokenType
	Value string
}
