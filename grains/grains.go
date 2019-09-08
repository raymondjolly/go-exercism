package grains

import (
	"github.com/pkg/errors"
)

//Square returns an error if the parameter is less than zero or
//greater than 64 OR it returns a square of number^(i-1).
func Square(number int) (uint64, error) {
	if number > 64 || number <= 0 {
		return 0, errors.New("incorrect range used")
	}
	return 1 << (uint(number) - 1), nil
}

// Total returns a value that is aways a constant at  2^64 - 1.
// Note that this is very confusing from the original documentation as
// that one would assume that the Total would take on the number of inputted
// squares which is not the case.
func Total() uint64 {
	return ^uint64(0)
}