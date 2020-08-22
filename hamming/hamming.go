package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return 0, errors.New("Error - Sequences have different lengths")
	}

	tam := len(a)

	distance := 0

	for i := 0; i < tam; i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
