package piscine

func BasicAtoi2(s string) int {
	var result int
	num := []rune(s)
	for i := 0; i < len(num); i++ {
		if '0' <= num[i] && num[i] <= '9' {
			result = result * 10 + int(num[i] - '0')
		} else {
			return 0
		}
	}
	return result
}