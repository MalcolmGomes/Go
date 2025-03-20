// Strings
package main

import (
	"fmt"
	"strings"
)

var pl = fmt.Println

func main() {
	sV1 := "A word"
	replacer := strings.NewReplacer("A", "Another")
	sV2 := replacer.Replace(sV1)
	pl(sV2)
	pl("Length:", len(sV2))
	pl("Contains Another :", strings.Contains(sV2, "Another"))
	pl("o index: ", strings.Index(sV2, "o"))
	pl("Replace :", strings.Replace(sV2, "o", "0", -1)) // -1 Searches everywhere. A 2 would replace first 2 matches.
	sV3 := "\nSome Words\n" // \ is a an escape sequence. Lets you use: \t Use \\ for backslash use in string
	pl(sV3)
	sV3 = strings.TrimSpace(sV3)
	pl(sV3)
	pl("Split: ", strings.Split("a-b-c-d", "-")) // Here - is a delimiter that is a character that appears between spaces of data.
	pl("Lower: ", strings.ToLower(sV2))
	pl("Upper: ", strings.ToUpper(sV2))
	pl("Prefix: ", strings.HasPrefix("tacocat", "taco"))
	pl("Suffix: ", strings.HasSuffix("tacocat", "cat"))
}