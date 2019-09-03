package bob

import (
	"regexp"
	"strings"
)
//Hey takes in an input and generates the moody responses that Bob should generate base on the following conditions:
//Bob answers 'Sure.' if you ask him a question.
//He answers 'Whoa, chill out!' if you yell at him.
//He answers 'Calm down, I know what I'm doing!' if you yell a question at him.
//He says 'Fine. Be that way!' if you address him without actually saying anything.
//He answers 'Whatever.' to anything else.
//Bob's conversational partner is a purist when it comes to written communication and
// always follows normal rules regarding sentence punctuation in English.

func Hey(greeting string) string {

	greeting = strings.TrimSpace(greeting)

	switch {
	case isSilent(greeting):
		return "Fine. Be that way!"
	case isQuestion(greeting) && isAllCaps(greeting) == true:
		return "Calm down, I know what I'm doing!"
	case isAllCaps(greeting)==true:
		return "Whoa, chill out!"
	case isQuestion(greeting)==true:
		return "Sure."
	default:
		return "Whatever."
	}

}

func isAllCaps(statement string) bool{
	re:= regexp.MustCompile("[a-zA-Z]")
	var capitalizedLetters = strings.ToUpper(statement)
	var containsLetter = re.Match([]byte(statement))

	return containsLetter && capitalizedLetters == statement
}

func isQuestion(statement string) bool{
	return strings.HasSuffix(statement, "?")
}

func isSilent(statement string) bool{
	return statement == ""
}
