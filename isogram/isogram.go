package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram (word string) bool{
	lettersInWord:=make(map[rune]bool, len(word))
	for _, letter:= range strings.ToLower(word){
		if _, encountered:= lettersInWord[letter];  encountered && unicode.IsLetter(letter){
			return false
		}
		lettersInWord[letter] = true
	}
	return true
}
