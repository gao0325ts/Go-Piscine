package piscine

func StrLen(s string) int {
	length := 0
	runes := []rune(s)
	for range runes {
		length++
	}
	return length
}
