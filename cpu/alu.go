package cpu

type ALU struct {
	a [8]bool
	b [8]bool
}

func (alu *ALU) Add() (result [8]bool, carryOut bool) {
	carry := false
	for i := 7; i >= 0; i-- {
		result[i], carry = fullAdder(alu.a[i], alu.b[i], carry)
	}

	return result, carry
}

func halfAdder(a, b bool) (sum, carry bool) {
	sum = Xor(a, b)
	carry = And(a, b)
	return sum, carry
}

func fullAdder(a, b, carryIn bool) (sum, carryOut bool) {
	sum1, carry1 := halfAdder(a, b)
	sum2, carry2 := halfAdder(sum1, carryIn)
	carryOut = Or(carry1, carry2)
	return sum2, carryOut
}
