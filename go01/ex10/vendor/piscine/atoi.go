package piscine

func Atoi(s string) int {
	var result int
	sign := 1
	start := 0
	num := []rune(s)
	if num[0] == '-' {
		sign = -1
		start++
	} else if num[0] == '+' {
		start++
	}
	for i := start; i < len(num); i++ {
		if '0' <= num[i] && num[i] <= '9' {
			result = result * 10 + int(num[i] - '0')
		} else {
			return 0
		}
	}
	return sign * result
}