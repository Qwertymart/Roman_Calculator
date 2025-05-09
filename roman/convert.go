package roman

import "fmt"

var romanToInt = map[string]int{
	"M": 1000, "CM": 900, "D": 500, "CD": 400,
	"C": 100, "XC": 90, "L": 50, "XL": 40,
	"X": 10, "IX": 9, "V": 5, "IV": 4, "I": 1,
}

func FromRoman(s string) (int, error) {
	var result int
	i := 0

	for i < len(s) {
		if i+1 < len(s) {
			if val, ok := romanToInt[s[i:i+2]]; ok {
				result += val
				i += 2
				continue
			}
		}
		if val, ok := romanToInt[s[i:i+2]]; ok {
			result += val
			i++
		} else {
			return 0, fmt.Errorf("invalid symbol %q", s[i])
		}
	}
	return result, nil
}

func ToRoman(n int) (string, error) {

	if n <= 0 || n > 3999 {
		return "", fmt.Errorf("invalid input: %d, must be between 1 and 3999", n)
	}
	var result string

	romanMap := []struct {
		Value  int
		Symbol string
	}{
		{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"},
		{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
		{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
	}

	for _, r := range romanMap {
		for n >= r.Value {
			result += r.Symbol
			n -= r.Value
		}
	}
	return result, nil
}
