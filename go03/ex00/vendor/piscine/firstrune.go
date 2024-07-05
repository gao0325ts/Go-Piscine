package piscine

func FirstRune(s string) rune {
	for _, r := range []rune(s) {
		return r
	}
	return 0
}