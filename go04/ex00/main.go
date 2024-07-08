package main

import (
	"os"
	"piscine"
)

func main() {
	name := os.Args[0]
	piscine.PrintProgramName(name)
}