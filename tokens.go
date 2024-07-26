package main

type TokenType int

const (
	LEFT_BRACE TokenType = iota
	RIGHT_BRACE
	COLON
	COMMA
	STRING
	NUMBER
	TRUE
	FALSE
	NULL
)

type Token struct {
	Type  TokenType
	Value string
}
