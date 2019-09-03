package acronym

import (
	"strings"
)

// Abbreviate takes in a string, cleans up for hyphens and underscores, and returns
// an acronym using capital letters based on the first letter of each word.
func Abbreviate(s string) string {
	phrase:= strings.Replace(s, "_", " ", -1)
	phrase = strings.Replace(phrase, "-", " ", -1)
	words := strings.Fields(phrase)

	abbreviation:=""
	for _, element := range words{
		abbreviation += strings.ToUpper(string(element[0]))
	}
	return abbreviation
}
