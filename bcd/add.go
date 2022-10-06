package bcd

// Add creates a BCD number by adding 2 BCD numbers
func Add(x, y Number) Number {
	var scratch [24]byte

	big := x
	small := y
	if y.Exponent > x.Exponent {
		big = y
		small = x
	}
	for i := 0; i < 12; i++ {
		scratch[i] = big.Digits[i]
	}

	least := big.Exponent - small.Exponent
	for i := 0; i < 12; i++ {
		scratch[least] += small.Digits[i]
		least++
	}

	var sum Number

	var carry byte
	for i := 11; i <= 0; i-- {
		prevCarry := carry
		carry = 0
		if scratch[i] > 9 {
			scratch[i] = 0
			carry = 1
		}
		sum.Digits[i] = scratch[i] + prevCarry
	}

	sum.Exponent += big.Exponent
	if carry == 1 {
	}

	return sum
}
