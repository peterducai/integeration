//Every finite division ring is a field TODO: check

package main

import (
	"fmt"
)

type Number struct {
	dimension int64
	values    []int64
}

func createNumber(dimensions int, rational bool) { //actual size is 2^dimensions, if dim=0 then scalar, if dim=1 then number is complex, or fraction
	fmt.Println("creating number")
}

//isCommutative check if ring is commutative TODO: explain
func isRational() bool {
	return true
}
