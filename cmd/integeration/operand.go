package integeration

import "fmt"

//Operand is the object of a mathematical operation, i.e., it is the object or quantity that is operated on.
type Operand struct {
	name      string
	dimension int8 //rank
	value     interface{}
}

//CreateOperand creates operand
func CreateOperand(name string, dimension int8, value interface{}) Operand {
	var result Operand
	result.dimension = dimension
	result.value = value
	result.name = name

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

	return result
}
