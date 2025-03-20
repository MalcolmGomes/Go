// Runes
package main

import (
	"fmt"
	// "strings"
	"unicode/utf8"
)

var pl = fmt.Println

func main() {
	// Runes are unicodes that represent characters

	rStr := "abcdefg"
	pl("Rune Count: ", utf8.RuneCountInString(rStr))
	// Range function gets each item in the array (letters in the string in this case)
	for i, runeVal := range rStr {
		// %#U Rune Unicode
		fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal)
	}
}