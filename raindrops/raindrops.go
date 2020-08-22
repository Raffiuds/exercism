package raindrops

import "fmt"

type raindrop struct {
	factor int
	sound  string
}

var raindrops = []raindrop{
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"},
}

func Convert(num int) string {

	sounds := ""

	for _, v := range raindrops {
		if num%v.factor == 0 {
			sounds += v.sound
		}
	}

	if sounds == "" {
		sounds = fmt.Sprintf("%d", num)
	}

	return sounds
}
