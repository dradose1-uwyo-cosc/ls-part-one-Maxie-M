// gols.go
// Maxie Machado
// ls-part-one
// February 26, 2026

package main

import (
	"bufio"
	"os"

	"lspt1/functions"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	args := os.Args[1:]

	useColor := functions.IsTerminal(os.Stdout)
	functions.SimpleLS(out, args, useColor)
}
