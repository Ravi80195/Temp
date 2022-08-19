
	//2. Do the same problem, but this time, implement the below code:
*/

package main

import (
	"fmt"
	"strings"
	"unicode"
)

type Interface interface {
	TransformRune(pos int)
	GetValueAsRuneSlice() []rune
}

type inputStruct struct {
	input     string
	position  int
	charIndex int
}

func (i *inputStruct) GetValueAsRuneSlice() []rune {
	return []rune(i.input)
}

func (i *inputStruct) TransformRune(pos int) {
	if unicode.IsLetter(rune(i.input[pos])) {
		if i.charIndex%i.position == 0 {
			i.input = i.input[:pos] + strings.ToUpper(string(i.input[pos])) + i.input[pos+1:]
		}
		i.charIndex++
	}
}

func MapString(i Interface) {
	for pos := range i.GetValueAsRuneSlice() {
		i.TransformRune(pos)
	}

}

func (i inputStruct) String() string {
	return fmt.Sprintf("%v", i.input)
}

func NewSkipString(pos int, input string) inputStruct {
	return inputStruct{input: strings.ToLower(input), position: pos, charIndex: 1}
}

// And change your code to look like this:
func main() {
	s := NewSkipString(3, "Aspiration.com")
	MapString(&s)
	fmt.Println(s)
}

/*
  Make sure your fmt.Println(s) output looks nice and clean, ie implement the Stringer interface.
*/
