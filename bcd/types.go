package bcd

import (
	"fmt"
	"strings"
	"unicode"
)

const significantDigits = 12

type Number struct {
	Sign     int8
	Exponent int8                    // signed 2s-complement
	Digits   [significantDigits]byte // most significant at index 0
}

// Implided decimal point between Digits[0] and Digits[1]

// Aton converts a suitably-formatted string to a BCD number type
func Aton(stringrep string) (Number, error) {
	stringrep = strings.TrimSpace(stringrep)

	var b Number

	sign := 0
	if stringrep[0] == '-' {
		sign = 1
		stringrep = stringrep[1:]
	}

	for stringrep[0] == '0' {
		stringrep = stringrep[1:]
	}

	digits := 0
	decimalPointAt := -1
	foundE := false
	eAt := -1

	for idx, r := range stringrep {
		if r == '.' {
			decimalPointAt = digits
		}
		if unicode.IsDigit(r) {
			if digits < significantDigits {
				b.Digits[digits] = byte(r - '0')
			}
			digits++
		}
		if r == 'e' || r == 'E' {
			foundE = true
			eAt = idx
			break
		}
	}

	var exponent int8
	if foundE {
		eAt++
		var exponentSign int8 = 1
		if stringrep[eAt] == '-' {
			exponentSign = -1
			eAt++
		}
		for _, r := range stringrep[eAt:] {
			exponent = 10*exponent + (int8(r) - '0')
		}
		exponent *= exponentSign
	}

	if decimalPointAt > -1 {
		exponent += int8(decimalPointAt - 1)
	} else {
		// implied decimal point after final digit
		exponent += int8(digits - 1)
	}

	b.Exponent = exponent
	b.Sign = int8(sign)

	return b, nil
}

// String converts a BCD number to a printable string
func (b Number) String() string {
	sign := ""
	if b.Sign != 0 {
		sign = "-"
	}

	stringrep := fmt.Sprintf("%s%d.", sign, b.Digits[0])

	var i int
	for i = 11; i > 0; i-- {
		if b.Digits[i] != 0 {
			i++
			break
		}
	}

	for j := 1; j < i; j++ {
		stringrep = fmt.Sprintf("%s%d", stringrep, b.Digits[j])
	}

	return fmt.Sprintf("%sE%d", stringrep, b.Exponent)
}
