package piscine

func IsUpper(s string) bool {
	for _, r := range []rune(s) {
		if !(r >= 'A' && r <= 'Z') {
			return false
		}
	}
	return true
}
