package utils

import "math/rand"

func GenerateInts(from, to, len int) []int {
	data := make([]int, len)
	diff := to - from + 1

	for i := 0; i < len; i++ {
		data[i] = rand.Intn(diff) + from
	}

	return data
}
