package piscine

func StrLen(s string) int {
	var length int
	for range []rune(s) {
		length++
	}
	return length
}

func StrRev(s string) string {
	runes := []rune(s)
	length := StrLen(s)
	for i := 0; i < length/2; i++ {
		runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
	}
	return string(runes)
}
