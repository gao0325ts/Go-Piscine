package piscine

import (
	_ "ft"
	_ "fmt"
)

func SortIntegerTable(table []int) {
	var sorted bool
	for sorted == false {
		sorted = true
		for i := 0; i < len(table) - 1; i++ {
			if table[i] > table[i + 1] {
				tmp := table[i]
				table[i] = table[i + 1]
				table[i + 1] = tmp
				sorted = false
			}
		}
	}
}