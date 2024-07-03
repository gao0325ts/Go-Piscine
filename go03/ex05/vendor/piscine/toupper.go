package piscine

func ToUpper(s string) string {
	runes := []rune(s)
	for i, r := range runes {
		if r >= 'a' && r <= 'z' {
			runes[i] -= 32
		}
	}
	return string(runes)
}