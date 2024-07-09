package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	dec := atoiBase(nbr, baseFrom)
	return itoaBase(dec, baseTo)
}

func strLen(s string) int {
	var length int
	for range s {
		length++
	}
	return length
}

func itoaBase(n int, base string) string {
	baseLen := strLen(base)
	var baseNum string
	if n == 0 {
		return string(base[0])
	}
	for n > 0 {
		baseNum = string(base[n%baseLen]) + baseNum
		n /= baseLen
	}
	return baseNum
}

func atoiBase(s, base string) int {
	baseLen := strLen(base)
	var decNum int
	for _, r := range s {
		for i, b := range base {
			if r == b {
				decNum = decNum*baseLen + i
			}
		}
	}
	return decNum
}
