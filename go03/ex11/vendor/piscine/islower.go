package piscine

func IsLower(s string) bool {
	for _, r := range []rune(s) {
		if !(r >= 'a' && r <= 'z') {
			return false
		}
	}
	return true
}
