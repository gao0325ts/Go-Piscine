package piscine

func StrLen(s string) int {
	var length int
	for range []rune(s) {
		length++
	}
	return length
}
