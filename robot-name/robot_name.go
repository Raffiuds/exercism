package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Robot string

var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var hashTable = make(map[string]bool)

func init() {
	rand.Seed(time.Now().UnixNano())

}

func (robot *Robot) Name() (string, error) {

	if len(hashTable) == 676000 {
		return "", errors.New("Exhausted NameSpace")
	}
	if *robot == "" {

		name := newName()

		for contains(name) {
			name = newName()
		}

		*robot = Robot(name)
		return name, nil
	} else {
		return string(*robot), nil
	}
}

func (robot *Robot) Reset() {

	*robot = ""
}

func newName() string {
	return fmt.Sprintf("%s%d%d%d",
		randStringRunes(2),
		randUInt(10),
		randUInt(10),
		randUInt(10))
}

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func randUInt(n int) int {
	return rand.Intn(n)
}

func contains(s string) bool {

	if _, ok := hashTable[s]; ok {
		return true
	}
	hashTable[s] = true

	return false
}
