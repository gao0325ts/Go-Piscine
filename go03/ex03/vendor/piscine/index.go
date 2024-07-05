package piscine

func strLen(s string) int {
	var length int
	for range []rune(s) {
		length++
	}
	return length
}

func Index(s string, toFind string) int {
	sLen := strLen(s)
	toFindLen := strLen(toFind)
	for i := 0; i <= sLen-toFindLen; i++ {
		found := true
		for j := 0; j < toFindLen; j++ {
			if s[i+j] != toFind[j] {
				found = false
				break
			}
		}
		if found {
			return i
		}
	}
	return -1
}
