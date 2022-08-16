
	//1 . Write a function with this signature:
	
  
package main

import (
	"fmt"
	"strings"
)

func CapitalizeEveryThirdAlphanumericChar(position int, input string) string {
	input = strings.ToLower(input)
	for pos := range []rune(input) {
		if pos%position == 0 && pos > 0 {
			input = input[:pos-1] + strings.ToUpper(string(input[pos-1])) + input[pos:]
		}
	}
	return input

}

// And change your code to look like this:
func main() {
	s := CapitalizeEveryThirdAlphanumericChar(3, "Aspiration.com")
	fmt.Println(s)
}


// 	that capitalizes *only* every third alphanumeric character of a string, so that if I hand you
// 	Aspiration.com
// 	You hand me back
// 	asPirAtiOn.cOm
// 	Please note: 
// 	- Characters other than each third should be downcased
// 	- For the purposes of counting and capitalizing every three characters, consider only alphanumeric
// 	  characters, ie [a-z][A-Z][0-9]. 
