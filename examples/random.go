// Random
package main

import (
	"fmt"
	"time"
	"math/rand"
)

var pl = fmt.Println

func main() {
	// Time is seconds from 1/1/70

	seedSecs := time.Now().Unix() // Seconds from 1970.
	rand.Seed(seedSecs) // Generating random seed using seconds from 1970
	randNum := rand.Intn(50) + 1 // Generate random values from 0 to 49
	pl("Random: ", randNum)

}