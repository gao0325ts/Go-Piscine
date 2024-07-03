package piscine

func LastRune(s string) rune {
	runes := []rune(s)
	var last rune
	for _, r := range runes {
		last = r
	}
	return last
}
