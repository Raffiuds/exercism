package luhn

import (
	"strconv"
	"strings"
)

func Valid(num string) bool {

	num = removeSpace(num)
	if len(num) <= 1 {
		return false
	}

	numbers, err := stringToIntArray(num)

	numbers = fistStep(numbers)
	if err != nil {
		return false
	}

	isValid := secondStep(numbers)

	return isValid
}

func removeSpace(s string) string {

	return strings.ReplaceAll(s, " ", "")
}

func stringToIntArray(s string) ([]int, error) {

	numbers := []int{}
	for i := len(s) - 1; i >= 0; i-- {
		i, err := strconv.Atoi(string(s[i]))
		if err != nil {
			return []int{}, err
		}
		numbers = append(numbers, i)
	}

	return numbers, nil
}

//The first step of the Luhn algorithm is to double every second digit,
//starting from the right. We will be doubling
//If doubling the number results in a number greater than 9 then subtract 9 from the product.
//The results of our doubling:
func fistStep(numbers []int) []int {

	double := make([]int, len(numbers))

	copy(double, numbers)

	for i := 0; i < len(double); i++ {
		if i%2 != 0 {
			double[i] = double[i] * 2
			if double[i] > 9 {
				double[i] = double[i] - 9
			}
		}
	}

	return double
}

//Then sum all of the digits
//If the sum is evenly divisible by 10, then the number is valid. This number is valid!

func secondStep(numbers []int) bool {

	sum := 0

	for _, num := range numbers {
		sum += num
	}

	return sum%10 == 0
}
