package piscine

func FirstRune(s string) rune {
	runes := []rune(s)
	for _, r := range runes {
		return r
	}
	return 0
}