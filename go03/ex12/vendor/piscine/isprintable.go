package piscine

func IsPrintable(s string) bool {
	for _, r := range []rune(s) {
		if !(r >= 32 && r <= 126) {
			return false
		}
	}
	return true
}
