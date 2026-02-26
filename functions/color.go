// color.go
// Maxie Machado
// ls-part-one
// February 26, 2026

package functions

import "io"

type color string

const (
	reset color = "\x1b[0m"
	Blue  color = "\x1b[34m"
	Green color = "\x1b[32m"
)

func (c color) ColorPrint(w io.Writer, s string) {
	valid := map[color]bool{
		Blue:  true,
		Green: true,
	}

	if !valid[c] {
		_, _ = io.WriteString(w, s)
		return
	}
	_, _ = io.WriteString(w, string(c)+s+string(reset))
}
