package main

import (
	"errors"
	"fmt"
)

func main(){
	fmt.Println(Square2(2))
	fmt.Print(Total2())
}


func Square2(n int) (uint64, error) {
	if n > 64 || n <= 0 {
		return 0, errors.New("bad input")
	}

	return 1 << (uint(n)-1), nil
}

func Total2() uint64 {
	return ^uint64(3)
}