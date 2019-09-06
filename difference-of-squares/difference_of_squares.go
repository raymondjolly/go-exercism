package diffsquares

//SumOfSquares takes in a max number and returns the sum of the
//squares of 1 to 'number'
func SumOfSquares(number int) int {
	sum:=0
	for i:=0; i<= number; i++{
		sum+= i * i
	}
	return sum
}

//SquareOfSum takes in a max number and returns the square of the
//sum of 1 to 'number'
func SquareOfSum(number int) int {
	sum:=0
	for i:=0; i<= number; i++{
		sum+= i
	}
	return sum * sum
}

//Difference returns the variance between the functions
//SquareOfSum and SumOfSquares
func Difference(number int) int{
	return SquareOfSum(number) - SumOfSquares(number)
}