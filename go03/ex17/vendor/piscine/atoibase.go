package piscine

func strLen(s string) int {
	var length int
	for range s {
		length++
	}
	return length
}

func isValidBase(base string) bool {
	if strLen(base) < 2 {
		return false
	}
	var extChars string
	for _, r := range base {
		if r == '-' || r == '+' {
			return false
		}
		for _, c := range extChars {
			if r == c {
				return false
			}
		}
		extChars += string(r)
	}
	return true
}

func Recursive(s, base string, len int) int {
	var result int
	for _, r := range s {
		for i, b := range base {
			if r == b {
				result = result * len + i
			}
		}
	}
	return result
}

func AtoiBase(s string, base string) int {
	if !isValidBase(base) {
		return 0
	}
	for _, r := range s {
		found := false
		for _, b := range base {
			if r == b {
				found = true
			}
		}
		if !found {
			return 0
		}
	}
	baseLen := strLen(base)
	return Recursive(s, base, baseLen)
}