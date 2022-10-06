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
	carry := byte(0)
	for i := 0; i < 12; i++ {
		scratch[least] += small.Digits[i] + carry
		carry = 0
		if scratch[least] > 9 {
			scratch[least] = 0
			carry = 1
		}
		least++
	}

	var sum Number

	var i int
	if carry == 1 {
		sum.Digits[0] = 1
		i = 1
		sum.Exponent = 1
	}
	for ; i < 12; i++ {
		sum.Digits[i] = scratch[i]
	}
	sum.Exponent += big.Exponent

	return sum
}
