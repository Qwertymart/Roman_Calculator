package roman

type Numeral struct {
	Roman string
	Value int
}

func NewRoman(s string) (*Numeral, error) {
	val, err := FromRoman(s)

	if err != nil {
		return nil, err
	}

	return &Numeral{s, val}, nil
}

func NewDecimal(n int) (*Numeral, error) {
	roman, err := ToRoman(n)

	if err != nil {
		return nil, err
	}
	return &Numeral{roman, n}, nil
}
