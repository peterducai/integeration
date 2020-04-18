package main

import (
	"fmt"

	"github.com/peterducai/integeration/models"
)

func main() {
	fmt.Println("integeration")

	s1 := models.CreateOperand("s1", 0, 5)
	fmt.Println(s1)

	//COMPLEX NUMBER
	complexnum := [...]int{10, 20}
	v1 := models.CreateOperand("v1", 1, complexnum)
	fmt.Println(v1)
	// matrix := integeration.CreateOperand(2, 4)
	// tensor := integeration.CreateOperand(0, 5)
}
