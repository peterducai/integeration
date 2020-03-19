package main

import (
	"fmt"

	"github.com/peterducai/integeration/cmd/integeration"
)

func main() {
	fmt.Println("integeration")

	scalar := integeration.CreateOperand(0, 5)
	fmt.Println(scalar)

	//COMPLEX NUMBER
	complexnum := [...]int{10, 20}
	vector := integeration.CreateOperand(1, complexnum)
	fmt.Println(vector)
	// matrix := integeration.CreateOperand(2, 4)
	// tensor := integeration.CreateOperand(0, 5)
}
