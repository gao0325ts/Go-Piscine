package main

import (
	"os"
	"piscine"
)

func main() {
	params := os.Args
	piscine.SortParams(params)
}