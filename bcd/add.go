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

	least := big.Exponent + small.Exponent
	for i := 0; i < 12; i++ {
		scratch[least] += small.Digits[i]
		least++
	}

	var sum Number
	for i := 0; i < 12; i++ {
		sum.Digits[i] = scratch[i]
	}
	sum.Exponent = big.Exponent

	return sum
}
