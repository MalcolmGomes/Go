// Reading string from standard input
package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
)

var pl = fmt.Println
// Comment
/* 
Multiline comment
*/

func main() {
	pl("What is your name?")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err == nil {
		pl("Hello", name)
	} else {
		log.Fatal(err)
	}
}