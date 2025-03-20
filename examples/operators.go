// Operators
package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	pl("5 + 4 = ", 5 + 4)
	pl("5 - 4 = ", 5 - 4)
	pl("5 / 4 = ", 5.0 / 4)
	pl("5 * 4 = ", 5 * 4)
	pl("5 % 4 = ", 5 % 4)

	mInt := 1
	mInt += mInt + 1 // Increment Operator
	mInt++ // Shorthand Increment
	
	pl("Float Precision =", 0.11111111111111111111111111111111 + 0.1111111111111111111111111111111111)
}