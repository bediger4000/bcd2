package bcd

const scratchSize = 24

// Add creates a BCD number by adding 2 BCD numbers
func Add(x, y Number) Number {
	var scratch [scratchSize]byte

	big := x
	small := y
	if y.Exponent > x.Exponent {
		big = y
		small = x
	}
	for i := 0; i < 12; i++ {
		scratch[i] = big.Digits[i]
	}

	// At what index into scratch[] should we start adding?
	index := big.Exponent - small.Exponent
	for i := 0; i < 12 && index < scratchSize; i++ {
		scratch[index] += small.Digits[i]
		index++
	}

	// Do carrying
	carry := byte(0)
	for i := 23; i >= 0; i-- {
		scratch[i] += carry
		carry = 0
		if scratch[i] > 9 {
			scratch[i] -= 10
			carry = 1
		}
	}

	var sum Number
	i := 0
	if carry == 1 {
		sum.Digits[0] = 1
		i = 1
		sum.Exponent = 1
	}

	for j := 0; j < 12 && i < 12; j++ {
		sum.Digits[i] = scratch[j]
		i++
	}

	sum.Exponent += big.Exponent

	return sum
}
