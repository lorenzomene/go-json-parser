package main

import "fmt"

type Parser struct {
	tokens  []Token
	current int
}

func NewParser(tokens []Token) *Parser {
	return &Parser{
		tokens:  tokens,
		current: 0,
	}
}

func (p *Parser) nextToken() Token {
	if p.current < len(p.tokens) {
		token := p.tokens[p.current]

		p.current++

		return token
	}
	return Token{Type: NULL, Value: ""}
}

func (p *Parser) peekToken() Token {
	if p.current < len(p.tokens) {
		return p.tokens[p.current]
	}
	return Token{Type: NULL, Value: ""}
}

func (p *Parser) Parse() (interface{}, error) {
	token := p.peekToken()

	switch token.Type {
	case LEFT_BRACE:
		p.parseObject()
	case LEFT_BRACKET:
		p.parseArray()
	}
	return nil, fmt.Errorf("unexpected token: %v, expected '{' or '['", token.Type)
}

func (p *Parser) parseArray() {

}

func (p *Parser) parseObject() {

}
