package piscine

func ToLower(s string) string {
	runes := []rune(s)
	for i, r := range runes {
		if r >= 'A' && r <= 'Z' {
			runes[i] += 32
		}
	}
	return string(runes)
}