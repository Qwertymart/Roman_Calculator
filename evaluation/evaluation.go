package evaluation

import (
	"fmt"
	"github.com/Qwertymart/Roman_Calculator/roman"
	"github.com/Qwertymart/Roman_Calculator/tokenizer"
)

func Evaluate(str string) (*roman.Numeral, error) {
	tokens, err := tokenizer.Tokenize(str)
	if err != nil {
		return nil, err
	}

	result, err := evaluateTokens(tokens)
	if err != nil {
		return nil, err
	}
	return result, err
}

func evaluateTokens(tokens []tokenizer.Token) (*roman.Numeral, error) {
	var operands []*roman.Numeral
	var operators []string

	precedence := map[string]int{
		"+": 1, "-": 1,
		"*": 2, "/": 2,
		"(": 0,
	}

	for _, tok := range tokens {
		switch tok.Type {

		case tokenizer.Number:
			{
				num, err := roman.FromRoman(tok.Value)
				if err != nil {
					return nil, err
				}
				operands = append(operands, &roman.Numeral{Roman: tok.Value, Value: num})
			}
		case tokenizer.Operator:
			{
				for len(operators) > 0 &&
					precedence[operators[len(operators)-1]] >= precedence[tok.Value] {
					operator := operators[len(operators)-1]
					operators = operators[:len(operators)-1]

					if err := applyOperator(&operands, operator); err != nil {
						return nil, err
					}
				}
				operators = append(operators, tok.Value)
			}
		case tokenizer.Left:
			{
				operators = append(operators, tok.Value)
			}

		case tokenizer.Right:
			{
				for len(operators) > 0 &&
					operators[len(operators)-1] != "(" {
					operator := operators[len(operators)-1]
					operators = operators[:len(operators)-1]

					if err := applyOperator(&operands, operator); err != nil {
						return nil, err
					}
				}

				if len(operators) == 0 || operators[len(operators)-1] != "(" {
					return nil, fmt.Errorf("mismatched parentheses")
				}

				operators = operators[:len(operators)-1] // delete (
			}
		}
	}
	for len(operators) > 0 {
		operator := operators[len(operators)-1]
		operators = operators[:len(operators)-1]
		if err := applyOperator(&operands, operator); err != nil {
			return nil, err
		}
	}

	if len(operands) != 1 {
		return nil, fmt.Errorf("invalid expression")
	}

	return operands[0], nil
}

func applyOperator(operands *[]*roman.Numeral, operator string) error {
	if len(*operands) < 2 {
		return fmt.Errorf("invalid expression")
	}

	right := (*operands)[len(*operands)-1]
	left := (*operands)[len(*operands)-2]

	*operands = (*operands)[:len(*operands)-2]

	var result *roman.Numeral
	var err error

	switch operator {

	case "+":
		result, err = left.Add(right)
	case "-":
		result, err = left.Subtract(right)
	case "*":
		result, err = left.Multiply(right)
	case "/":
		result, err = left.Divide(right)
	default:
		return fmt.Errorf("unsupported operator: %s", operator)
	}

	if err != nil {
		return err
	}

	*operands = append(*operands, result)
	return nil
}
