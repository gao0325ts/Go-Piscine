package piscine

func BasicAtoi2(s string) int {
	var n int
	for _, r := range []rune(s) {
		if '0' <= r && r <= '9' {
			n = n*10 + int(r-'0')
		} else {
			return 0
		}
	}
	return n
}
