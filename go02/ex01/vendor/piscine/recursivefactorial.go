package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 {
		return 0
	} else if nb == 0 {
		return 1
	}
	next := RecursiveFactorial(nb - 1)
	maxInt := int(^uint(0) >> 1)
	if next > 0 && nb > maxInt/next {
		return 0
	}
	return nb * next
}
