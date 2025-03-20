// Casting
package main

import (
	"fmt"
	"strconv"
	"reflect"
)

var pl = fmt.Println

func main() {
	cV1 := 1.5
	cV2 := int(cV1)
	pl(cV2)

	cV3 := "500000000"
	cV4, err := strconv.Atoi(cV3)
	pl(cV4, err, reflect.TypeOf(cV4))

	cV5 := 5000000
	cV6 := strconv.Itoa(cV5)
	pl(cV6)

	cV7 := "3.14"
	if cV8, err := strconv.ParseFloat(cV7, 64); err == nil {
		pl(cV8)
	}

	cV9 := fmt.Sprintf("%f", 3.14)
	pl(cV9)

}