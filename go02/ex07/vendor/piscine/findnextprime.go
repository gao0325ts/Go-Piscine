package piscine

func FindNextPrime(nb int) int {
	var IsPrime func(nb int) bool
	IsPrime = func(nb int) bool {
		if nb <= 1 {
			return false
		} else if nb <= 3 {
			return true
		}
		for i := 2; i <= nb/2; i++ {
			if nb%i == 0 {
				return false
			}
		}
		return true
	}
	if nb <= 2 {
		return 2
	}
	for ; ; nb++ {
		if IsPrime(nb) {
			return nb
		}
	}
	return nb
}
