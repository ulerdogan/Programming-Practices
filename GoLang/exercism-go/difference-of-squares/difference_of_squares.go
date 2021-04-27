package diffsquares

func SquareOfSum(num int) (result int) {

	for i := 0; i <= num; i++ {
		result += i
	}

	return result * result
}

func SumOfSquares(num int) (result int) {

	for i := 0; i <= num; i++ {
		result += i * i
	}

	return result
}

func Difference(num int) (difference int) {

	return SquareOfSum(num) - SumOfSquares(num)
}
