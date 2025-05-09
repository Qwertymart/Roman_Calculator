package roman

import "fmt"

func (n *Numeral) Add(other *Numeral) (*Numeral, error) {
	sum := n.Value + other.Value
	romanStr, err := ToRoman(sum)
	if err != nil {
		return nil, err
	}
	return &Numeral{romanStr, sum}, nil
}

func (n *Numeral) Subtract(other *Numeral) (*Numeral, error) {
	diff := n.Value - other.Value
	if diff <= 0 {
		return nil, fmt.Errorf("result out of range")
	}
	romanStr, err := ToRoman(diff)
	if err != nil {
		return nil, err
	}
	return &Numeral{romanStr, diff}, nil
}

func (n *Numeral) Multiply(other *Numeral) (*Numeral, error) {
	product := n.Value * other.Value
	romanStr, err := ToRoman(product)
	if err != nil {
		return nil, err
	}
	return &Numeral{romanStr, product}, nil
}

func (n *Numeral) Divide(other *Numeral) (*Numeral, error) {
	if other.Value == 0 {
		return nil, fmt.Errorf("division by zero")
	}
	quotient := n.Value / other.Value
	if quotient == 0 {
		return nil, fmt.Errorf("quotient is zero")
	}
	romanStr, err := ToRoman(quotient)
	if err != nil {
		return nil, err
	}
	return &Numeral{romanStr, quotient}, nil
}
