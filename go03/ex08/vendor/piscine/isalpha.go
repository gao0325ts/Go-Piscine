package piscine

func IsAlpha(s string) bool {
	for _, r := range []rune(s) {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			continue
		} else {
			return false
		}
	}
	return true
}
