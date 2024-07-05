package piscine

func NRune(s string, n int) rune {
	for i, r := range []rune(s) {
		if i == n {
			return r
		}
	}
	return 0
}