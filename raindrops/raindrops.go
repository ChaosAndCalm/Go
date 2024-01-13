package raindrops

import (
	"strconv"
)

func Convert(number int) string {

	var res string
	check := false

	if number%3 == 0 {
		res += "Pling"
		check = true
	}

	if number%5 == 0 {
		res += "Plang"
		check = true
	}

	if number%7 == 0 {
		res += "Plong"
		check = true
	}

	if !check {
		res = strconv.Itoa(number)
	}

	return res
}
