package leap

// IsLeapYear is a function which calculates if a given year is a leap year.
// Check https://www.livescience.com/45768-gregorian-calendar.html for a historical
// guidance
func IsLeapYear(year int) bool {
	if year % 400 ==0{
		return true
	} else if year %100 ==0{
		return false
	} else if year %4==0{
		return true
	} else {
		return false
	}
}
