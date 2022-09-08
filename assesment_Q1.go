//1 . Write a function with this signature:

package main

import (
	"fmt"
	"regexp"
	"strings"
)

// Below method would transform the the input rune to the expected output.
// It has logic implemented where we have logic to check if it is a unicode character and if it is a unicode character then we check for the position.
// We check for the position for every third character ignoring the special characters in previous condition.
// If the 3 position is found we capitalize the character at that particular position.
// We do it in a similar way for all the characters.
func CapitalizeEveryThirdAlphanumericChar(position int, input string) string {
	input = strings.ToLower(input)
	charIndex := 1
	re := regexp.MustCompile("^[a-z0-9]*$")
	for pos := range []rune(input) {
		if re.MatchString(string(input[pos])) {
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
	s := CapitalizeEveryThirdAlphanumericChar(3, "As3piration.com")
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
