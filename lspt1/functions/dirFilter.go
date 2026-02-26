// dirFilter.go
// Maxie Machado
// ls-part-one
// February 26, 2026

package functions

import "os"

func dirFilter(entries []os.DirEntry) []os.DirEntry {
	out := make([]os.DirEntry, 0, len(entries))
	for _, e := range entries {
		name := e.Name()
		if len(name) == 0 || name[0] == '.' {
			continue
		}
		out = append(out, e)
	}
	return out
}
