package piscine

func NRune(s string, n int) rune {
	runes := []rune(s)
	for i, r := range runes {
		if i == n {
			return r
		}
	}
	return 0
}