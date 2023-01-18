package bcd

// Sub creates a BCD number by subtracting one BCD number from another
func Sub(x, y Number) Number {

	// (+x) - (-y)
	if x.Sign == 0 && y.Sign == 1 {
		y.Sign = 0
		return Add(x, y)
	}

	// (-x) - (-y)
	if x.Sign == 1 && y.Sign == 1 {
		x.Sign = 0
		y.Sign = 0
		answer := Add(x, y)
		answer.Sign = 1
		return answer
	}

	// Either (+x) - (+y) or (-x) - (+y)

	var sum Number

	return sum
}
