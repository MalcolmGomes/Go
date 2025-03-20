// Conditional Statements
package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	// Conditional Operators: > < >= <= == !=
	// Logical Operators: && || !

	iAge := 8
	if (iAge >= 1) && (iAge <= 18) {
		pl("Important Birthday")
	} else if iAge == 21 || iAge == 50 {
		pl("Important Birthday")
	} else if iAge >= 65 {
		pl("Important Birthday") 
	} else {
		pl("Not an Important Birthday")
	}

	// bool true or false
	pl("!true =", !true)
}