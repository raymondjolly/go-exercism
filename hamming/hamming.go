package hamming

import "errors"


//Distance takes in two equal length strands of DNA and will return the number of
//mutations between the two.
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return 0, errors.New("variable strands must be of equal length")
	}

	differenceCount := 0
	for i:= range a{
		if a[i] != b[i]{
			differenceCount++
		}
	}
	return differenceCount, nil
}
