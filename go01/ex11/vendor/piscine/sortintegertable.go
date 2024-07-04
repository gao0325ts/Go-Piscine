package piscine

func SortIntegerTable(table []int) {
	var sorted bool
	for sorted == false {
		sorted = true
		for i := range table {
			if i == 0 {
				continue
			}
			if table[i-1] > table[i] {
				table[i-1], table[i] = table[i], table[i-1]
				sorted = false
			}
		}
	}
}
