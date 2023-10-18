package utils

import "strconv"

func StringToInt(str string) int {
	// string to int
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}
