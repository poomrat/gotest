package fizzbuzz

import "strconv"

func Say(n int) string {
	if n == 3 {
		return "fizz"
	}
	return strconv.Itoa(n)
}
