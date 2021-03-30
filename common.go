package pdputils

import (
	"log"
	"strconv"
)

func getZeroIndices(arr []int) []int {
	var indices []int
	for i, elem := range arr {
		if elem == -1 {
			indices = append(indices, i)
		}
	}
	return indices
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func unsafeParse(number string) int {
	parsed, err := strconv.Atoi(number)
	check(err)
	return parsed
}
