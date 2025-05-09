package tokenizer

import (
	"fmt"
	"strings"
	"unicode"
)

type TokenType int

const (
	Number TokenType = iota
	Operator
	Left
	Right
)

type Token struct {
	Type  TokenType
	Value string
}

func Tokenize(str string) ([]Token, error) {
	str = strings.ReplaceAll(str, " ", "")
	var tokens []Token
	i := 0
	for i < len(str) {
		c := str[i]
		switch {
		case isRoman(c):
			start := i
			for i < len(str) && isRoman(str[i]) {
				i++
			}
			romanStr := str[start:i]

			tokens = append(tokens, Token{
				Type:  Number,
				Value: romanStr,
			})

		case c == '+' || c == '-' || c == '*' || c == '/':
			tokens = append(tokens, Token{
				Type:  Operator,
				Value: string(c),
			})
			i++

		case c == '(':
			tokens = append(tokens, Token{
				Type:  Left,
				Value: string(c),
			})
			i++

		case c == ')':
			tokens = append(tokens, Token{
				Type:  Right,
				Value: string(c),
			})
			i++

		default:
			return nil, fmt.Errorf("unexpected character: %c", c)
		}
	}
	return tokens, nil
}

func isRoman(c byte) bool {
	return strings.ContainsRune("IVXLCDM", rune(unicode.ToUpper(rune(c))))
}
