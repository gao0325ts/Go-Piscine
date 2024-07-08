package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	capNext := true
	for i, r := range runes {
		if r >= 'a' && r <= 'z' {
			if capNext {
				runes[i] -= ('a' - 'A')
			}
			capNext = false
		} else if r >= 'A' && r <= 'Z' {
			if !capNext {
				runes[i] += ('a' - 'A')
			}
			capNext = false
		} else if r >= '0' && r <= '9' {
			capNext = false
		} else {
			capNext = true
		}
	}
	return string(runes)
}
