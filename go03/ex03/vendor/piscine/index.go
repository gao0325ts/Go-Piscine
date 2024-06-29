package piscine

func Index(s string, toFind string) int {
	runes := []rune(s)
	toFindRunes := []rune(toFind)
	for i := 0; i <= len(runes)-len(toFindRunes); i++ {
		found := true
		for j := 0; j < len(toFindRunes); j++ {
			if runes[i+j] != toFindRunes[j] {
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
