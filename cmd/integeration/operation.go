package integeration

// Operation is a function which takes zero or more input values (called operands) to a well-defined output value
// There are two common types of operations: unary and binary.
// Unary operations involve only one value, such as negation and trigonometric functions.
// Binary operations, on the other hand, take two values, and include addition,
// subtraction, multiplication, division, and exponentiation.
type Operation struct {
	symbol string
	name   string
	optype string //unary or binary
}

//UNARY OPERATION

// Negation or additive inverse of a number a is the number that, when added to a, yields zero.
// For a number and, generally, in any ring, the additive inverse can be calculated using
// multiplication by −1; that is, −n = −1 × n . Examples of rings of numbers are integers,
// rational numbers, real numbers, and complex numbers.
func Negation() {

}

//BINARY OPERATION

// Addition +, plus (addition)
func Addition() {

}

// Subtraction −, minus (subtraction)
func Subtraction() {

}

// Division ÷, obelus (division)
func Division() {

}

// ×, times (multiplication)
