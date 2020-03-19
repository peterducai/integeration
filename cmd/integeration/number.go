//Every finite division ring is a field TODO: check

package integeration

import (
	"fmt"
)

type isRational struct {
}

//Number
//constructed from ratios (or fractions) of integers.
//the irrational numbers are all the real numbers which are not rational numbers
//
//A rational number is any number that can be expressed as the quotient or fraction p/q of two integers, a numerator p and a non-zero denominator q.
//Since q may be equal to 1, every integer is a rational number. The set of all rational numbers, often referred to as "the rationals", the field of
// rationals or the field of rational numbers is usually denoted by a boldface Q
//
//Among irrational numbers are the ratio π of a circle's circumference to its diameter, Euler's number e, the golden ratio φ, and the square root of two
//
//As a consequence of Cantor's proof that the real numbers are uncountable and the rationals countable, it follows that almost all real numbers are irrational.
type Number struct {
	dimension  int64
	values     []int64
	irrational bool
}

//CreateNumber creates number. number has n dimensions and can be either rational or irrational
func CreateNumber(n int, rational bool) { //actual size is 2^dimensions, if dim=0 then scalar, if dim=1 then number is complex, or fraction
	fmt.Println("creating number")
}
