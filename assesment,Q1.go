
	//1 . Write a function with this signature:
	
  
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func CapitalizeEveryThirdAlphanumericChar(position int, input string) string {
	input = strings.ToLower(input)
	charIndex := 1
	for pos := range []rune(input) {
		if unicode.IsLetter(rune(input[pos])) {
			if charIndex%position == 0 {
				input = input[:pos] + strings.ToUpper(string(input[pos])) + input[pos+1:]
			}
			charIndex++
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
