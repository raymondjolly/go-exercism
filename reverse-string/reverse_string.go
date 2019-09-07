package reverse

import "unicode/utf8"

//This is not has easy as one typically thinks as the test cases
//include UTF-8 symbols and thus the function must take these characters
//into account.  Runes in Go can help in this situation.


//Revere takes in a string and inverses the location of each character
//Test cases include UTF characters which is wny runes must be used.
//Rune is character literal which represents int32 value and each int
// value is mapped to Unicode code point https://en.wikipedia.org/wiki/Code_point
//It represents any characters of any alphabet from any language in the world.

func Reverse(word string) string{

	newWordMap:= make([]rune, utf8.RuneCountInString(word))

	i:= len(newWordMap)

	for _, character := range word{
		i--
		newWordMap[i] = character
	}
	return string(newWordMap)
}