package grains

import (
	"errors"
	"math"
)

const totalGrains = 18446744073709551615

func Square(square int) (uint64, error) {

	if !isValid(square) {
		return 0, errors.New("Error: input between 1 and 64")
	}

	return uint64(math.Exp2(float64((square - 1)))), nil
}

func Total() uint64 {
	return totalGrains
}

func isValid(i int) bool {
	return i >= 1 && i <= 64
}
