package integeration

import "fmt"

//Operand is the object of a mathematical operation, i.e., it is the object or quantity that is operated on.
type Operand struct {
	dimension int64 //rank
}

//CreateOperand creates operand
func CreateOperand(dimension int) {
	switch dimension {
	case 0:
		fmt.Println("creating Scalar")
	case 1:
		fmt.Println("creating Vector")
	case 2:
		fmt.Println("creating Matrix")
	case 4:
		fmt.Println("creating 3D tensor")
	default:
		fmt.Println("High dimensional Operands are NOT SUPPORTED")
	}
}
