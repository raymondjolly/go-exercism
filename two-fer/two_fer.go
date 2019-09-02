// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.


package twofer

import "fmt"

// ShareWith will take a value as a string and return a string
// value of "One for [name], one for me". If name == "", the default
// value will be "you".
func ShareWith(name string) string {

	if name == "" {
		return "One for you, one for me."
	} else {
		return fmt.Sprintf("One for %v, one for me.", name)
	}
}
