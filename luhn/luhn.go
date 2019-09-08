package luhn

import (
	"unicode"
	"unicode/utf8"
)

func Valid(card string) bool{
	hasLetter(card)
	newCard:=cleanNumbers(card)
	newCard = luhnProcedure(newCard)
	//return luhnCheck(newCard)
}

//hasLetter returns a boolean to determine if a card has
//letters in the string which are illegal
func hasLetter(s string) bool{

	for _, value :=range s{
		if !unicode.IsLetter(value){
			return false
		}
	}
	return true
}

//cleanNumbers will take in a string and return a string
//that cleans out any spaces, dashes, or underscores
func cleanNumbers(s string) string {
	newCard:=""

	for _, value := range s{
		if unicode.IsDigit(value){
			newCard+=string(value)
		}
	}
	return newCard
}

//luhnProcedure is a function which takes in a string and
//returns a string where the numeric value second from the right
//is doubled
func luhnProcedure(s string) string{
	newCard:=make([]rune, utf8.RuneCountInString(s))
	length:= len(s)

	for index, value:= range s{
		newCard[index] =value
	}

	for i:=len(newCard)-1; 0<=i; i--{
		if newCard[length-2] >9{
			newCard[i] = newCard[length-2] -9
		} else{
			i = length -2
			newCard[i] = newCard[i] *2
		}
	}
	//
	//for index, value:=range newCard{
	//	if newCard[length -2] > 9 {
	//		newCard[index] = newCard[length-2] - 9
	//	} else {
	//		index=length - 2
	//		newCard[index] = value * 2
	//	}
	//}
	return string(newCard)
}

//func luhnCheck(s string) bool{
//	sum:=0
//	for _, value:=range s{
//		sum+=int(value)
//	}
//	if sum % 10 == 0 {
//		return true
//	}
//	return false
//}





