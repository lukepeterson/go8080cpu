package cpu

// And implements an AND logic gate
//
// A B | O
// ----+--
// 0 0 | 0
// 0 1 | 0
// 1 0 | 0
// 1 1 | 1
func And(a, b bool) bool {
	return a && b
}

// Or implements an OR logic gate
//
// A B | O
// ----+--
// 0 0 | 0
// 0 1 | 1
// 1 0 | 1
// 1 1 | 1
func Or(a, b bool) bool {
	return a || b
}

// Not implements a NOT logic gate
//
// A | O
// --+--
// 0 | 1
// 1 | 0
func Not(a bool) bool {
	return !a
}

// Nand implements a NAND logic gate
//
// A B | O
// ----+--
// 0 0 | 1
// 0 1 | 1
// 1 0 | 1
// 1 1 | 0
func Nand(a, b bool) bool {
	return Not(And(a, b))
}

// Xor implements an XOR logic gate
//
// A B | O
// ----+--
// 0 0 | 0
// 0 1 | 1
// 1 0 | 1
// 1 1 | 0
func Xor(a, b bool) bool {
	return a != b
}
