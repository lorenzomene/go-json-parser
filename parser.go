package main

import (
	"fmt"
	"strconv"
)

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
		return p.parseObject()
	case LEFT_BRACKET:
		return p.parseArray()
	default:
		return nil, fmt.Errorf("unexpected token: %v, expected '{' or '['", token.Type)
	}
}

func (p *Parser) parseArray() ([]interface{}, error) {
	arr := []interface{}{}

	if p.nextToken().Type != LEFT_BRACKET {
		return nil, fmt.Errorf("unexpected token, expected '[ (LEFT_BRACKET) at the start of an array")
	}

	for {
		token := p.peekToken()

		if token.Type == RIGHT_BRACKET {
			p.nextToken()
			break
		}

		value, err := p.parseValue()
		if err != nil {
			return nil, err
		}

		arr = append(arr, value)

		token = p.peekToken()
		if token.Type == COMMA {
			p.nextToken()
		} else if token.Type != RIGHT_BRACKET {
			return nil, fmt.Errorf("expected COMMA or RIGHT_BRACKET, got %v", token.Type)
		}
	}

	return arr, nil
}

func (p *Parser) parseObject() (map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if p.nextToken().Type != LEFT_BRACE {
		return nil, fmt.Errorf("unexpected token, expected '{' (LEFT_BRACE) at the start of an object")
	}

	for {
		token := p.peekToken()

		if token.Type == RIGHT_BRACE {
			p.nextToken()
			break
		}

		if token.Type != STRING {
			return nil, fmt.Errorf("unexpected token type: %v, expected STRING", token.Type)
		}
		key := token.Value
		p.nextToken()

		if p.nextToken().Type != COLON {
			return nil, fmt.Errorf("unexpected token type: %v, expected COLON after key: %v", token.Type, key)
		}

		value, err := p.parseValue()

		if err != nil {
			return nil, err
		}
		obj[key] = value

		token = p.peekToken()
		if token.Type == COMMA {
			p.nextToken()
		} else if token.Type != RIGHT_BRACE {
			return nil, fmt.Errorf("expected COMMA or RIGHT_BRACE after value for key %v", key)
		}
	}

	return obj, nil
}

func (p *Parser) parseValue() (interface{}, error) {
	token := p.peekToken()

	switch token.Type {
	case STRING:
		p.nextToken()
		return token.Value, nil
	case NUMBER:
		p.nextToken()
		num, err := strconv.ParseFloat(token.Value, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid number format: %v", token.Value)
		}
		return num, nil
	case TRUE:
		p.nextToken()
		return true, nil
	case FALSE:
		p.nextToken()
		return false, nil
	case NULL:
		p.nextToken()
		return nil, nil
	case LEFT_BRACE:
		return p.parseObject()
	case LEFT_BRACKET:
		return p.parseArray()
	default:
		return nil, fmt.Errorf("unexpected token: %v", token.Type)
	}
}
