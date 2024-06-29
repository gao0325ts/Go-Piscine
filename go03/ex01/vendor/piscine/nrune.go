package piscine

func NRune(s string, n int) rune {
	runes := []rune(s)
	if n < 0 || n > len(runes) {
		return 0
	}
	return runes[n-1]
}

// FIXME
// goroutine 1 [running]:
// piscine.NRune(...)
// /Users/stakada/Desktop/go-pis-gh/go03/ex01/vendor/piscine/nrune.go:12
// main.main()
// /Users/stakada/Desktop/go-pis-gh/go03/ex01/main.go:13 +0x145
// exit status 2