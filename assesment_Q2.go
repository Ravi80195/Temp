//2. Do the same problem, but this time, implement the below code:

package main

import (
	"fmt"
	"regexp"
	"strings"
)

// Below is of type Interface with Transform and GetValueRuneSlice Methods
type Interface interface {
	TransformRune(pos int)
	GetValueAsRuneSlice() []rune
}

// InputStruct which have input, position and charIndex as the parameters
type inputStruct struct {
	input     string
	position  int
	charIndex int
}

// The below method would return the input value as rune slice
func (i *inputStruct) GetValueAsRuneSlice() []rune {
	return []rune(i.input)
}

// Below method would transform the the input rune to the expected output.
// It has logic implemented where we have logic to check if it is a unicode character and if it is a unicode character then we check for the position.
// We check for the position for every third character ignoring the special characters in previous condition.
// If the 3 position is found we capitalize the character at that particular position.
// We do it in a similar way for all the characters.
func (i *inputStruct) TransformRune(pos int) {
	re := regexp.MustCompile("^[a-z0-9]*$")
	if re.MatchString(string(i.input[pos])) {
		if i.charIndex%i.position == 0 {
			i.input = i.input[:pos] + strings.ToUpper(string(i.input[pos])) + i.input[pos+1:]
		}
		i.charIndex++
	}
}

// MapString function would accept interface as a parameter and which does both the get Rune Slice and Transform the Rune to the expected output.
func MapString(i Interface) {
	for pos := range i.GetValueAsRuneSlice() {
		i.TransformRune(pos)
	}

}

// The String method is used to format the obtained result.
func (i inputStruct) String() string {
	return fmt.Sprintf("%v", i.input)
}

// NewSkipString would take position and input string as input and returns the input struct as output.
func NewSkipString(pos int, input string) inputStruct {
	return inputStruct{input: strings.ToLower(input), position: pos, charIndex: 1}
}

// And change your code to look like this:
func main() {
	s := NewSkipString(3, "As3piration.com")
	MapString(&s)
	fmt.Println(s)
}

/*
  Make sure your fmt.Println(s) output looks nice and clean, ie implement the Stringer interface.
*/
