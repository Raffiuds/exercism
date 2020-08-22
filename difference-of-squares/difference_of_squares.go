package diffsquares

func SquareOfSum(num int) int {

	sum := 0

	for i := 1; i <= num; i++ {
		sum += i
	}
	return sum * sum
}

func SumOfSquares(num int) int {

	sum := 0

	for i := 1; i <= num; i++ {
		sum += i * i
	}
	return sum

}

func Difference(num int) int {

	return SquareOfSum(num) - SumOfSquares(num)
}
