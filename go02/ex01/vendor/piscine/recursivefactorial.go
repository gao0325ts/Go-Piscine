package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 {
		return 0
	} else if nb == 0 {
		return 1
	}
	maxInt := int(^uint(0) >> 1)
	prev := RecursiveFactorial(nb - 1)
	if prev > 0 && nb > maxInt/prev {
		return 0
	}
	return nb * prev
}
