package gates

// And implements an AND logic gate
//
// A B OUT
// 0 0 0
// 0 1 0
// 1 0 0
// 1 1 1
func And(a bool, b bool) bool {
	return a && b
}

// Or implements an OR logic gate
//
// A B OUT
// 0 0 0
// 0 1 1
// 1 0 1
// 1 1 1
func Or(a bool, b bool) bool {
	return a || b
}

// Not implements a NOT logic gate
//
// A OUT
// 0 1
// 1 0
func Not(a bool) bool {
	return !a
}

// Nand implements a NAND logic gate
//
// A B OUT
// 0 0 1
// 0 1 1
// 1 0 1
// 1 1 0
func Nand(a bool, b bool) bool {
	return Not(And(a, b))
}
