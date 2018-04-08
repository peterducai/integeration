//Every finite division ring is a field TODO: check

package main

import (
	"fmt"
)

type Number struct {
	dimension int64
	values    []int64
}

//createNumber creates number. number has n dimensions and can be either rational or irrational
func createNumber(n int, rational bool) { //actual size is 2^dimensions, if dim=0 then scalar, if dim=1 then number is complex, or fraction
	fmt.Println("creating number")
}

//isCommutative check if ring is commutative TODO: explain
func isRational() bool {
	return true
}
