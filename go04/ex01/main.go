package main

import (
	"os"
	"piscine"
)

func main() {
	params := os.Args
	piscine.PrintParams(params)
}