package Helper

import (
	"time"

	"math/rand"
)

func GenerateRandomNumber(length int) int {
	if length <= 0 {
		return 0
	}

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	min := 1
	for i := 1; i < length; i++ {
		min *= 10
	}
	max := min*10 - 1

	return r.Intn(max-min+1) + min
}
