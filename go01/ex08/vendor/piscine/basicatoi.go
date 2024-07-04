package piscine

func StrLen(s string) int {
	var length int
	for range []rune(s) {
		length++
	}
	return length
}

func BasicAtoi(s string) int {
	var n int
	for _, r := range []rune(s) {
		if '0' <= r && r <= '9' {
			n = n*10 + int(r-'0')
		}
	}
	return n
}
