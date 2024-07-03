package piscine

func Index(s string, toFind string) int {
	runes := []rune(s)
	toFindRunes := []rune(toFind)
	for i, _ := range runes {
		found := true
		for j, _ := range toFindRunes {
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
