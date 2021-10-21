package random

import (
	"math/rand"
)

func GetRandomPort() int {
	port := 12306
	port += rand.Int() % 44444
	return port
}
