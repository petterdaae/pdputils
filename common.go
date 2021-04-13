package pdputils

import (
	"log"
	"strconv"
)

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
