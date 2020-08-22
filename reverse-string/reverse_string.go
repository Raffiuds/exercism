package reverse

func Reverse(word string) string {

	var reverse string

	for _, s := range word {
		reverse = string(s) + reverse
	}

	return reverse
}
