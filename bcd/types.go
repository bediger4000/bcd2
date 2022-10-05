package bcd

import (
	"fmt"
	"strings"
	"unicode"
)

type Number struct {
	Sign     byte
	Exponent byte
	Digits   [12]byte
}

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
			if digits < 12 {
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

	var exponent byte
	if foundE {
		for _, r := range stringrep[eAt+1:] {
			exponent = 10*exponent + (byte(r) - '0')
		}
	}

	if decimalPointAt > -1 {
		exponent += byte(decimalPointAt - 1)
	}

	b.Exponent = exponent
	b.Sign = byte(sign)

	return b, nil
}

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
