package piscine

func Atoi(s string) int {
	sign := 1
	var n int
	for i, r := range []rune(s) {
		if i == 0 {
			if r == '-' {
				sign = -1
				continue
			} else if r == '+' {
				continue
			}
		}
		if '0' <= r && r <= '9' {
			n = n*10 + int(r-'0')
		} else {
			return 0
		}
	}
	return sign * n
}
