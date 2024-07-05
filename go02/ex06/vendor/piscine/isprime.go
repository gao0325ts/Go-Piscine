package piscine

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	} else if nb <= 3 {
		return true
	} else if nb%2 == 0 || nb%3 == 0 {
		return false
	}
	for i := 5; i <= nb/i; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
