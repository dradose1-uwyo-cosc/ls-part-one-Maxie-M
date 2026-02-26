// simplels.go
// Maxie Machado
// ls-part-one
// February 25, 2026

package functions

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
)

func SimpleLS(w io.Writer, args []string, useColor bool) {
	if len(args) == 0 {
		printDirListing(w, ".", useColor)
		return
	}
	files, dirs := collectTargets(args)

	sort.Strings(files)
	sort.Strings(dirs)

	for _, p := range files {
		printSingleTarget(w, p, useColor)
	}

	if len(dirs) == 0 {
		return
	}

	multiDir := len(dirs) > 1
	for i, d := range dirs {
		if multiDir {
			fmt.Fprintf(w, "%s:\n", d)
		}
		printDirListing(w, d, useColor)

		if multiDir && i != len(dirs)-1 {
			fmt.Fprintln(w)
		}
	}
}

func collectTargets(args []string) ([]string, []string) {
	files := make([]string, 0, len(args))
	dirs := make([]string, 0, len(args))

	for _, p := range args {
		info, err := os.Lstat(p)
		if err != nil {
			fmt.Fprintf(os.Stderr, "gols: %s: %v\n", p, err)
			continue
		}
		if info.IsDir() {
			dirs = append(dirs, p)
		} else {
			files = append(files, p)
		}
	}
	return files, dirs
}

func printSingleTarget(w io.Writer, path string, useColor bool) {
	info, err := os.Lstat(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "gols: %s: %v\n", path, err)
		return
	}
	name := filepath.Base(path)

	if useColor && isExecutableReg(info.Mode()) {
		Green.ColorPrint(w, name)
		fmt.Fprintln(w)
		return
	}
	fmt.Fprintln(w, name)
}

func printDirListing(w io.Writer, dir string, useColor bool) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "gols: %s: %v\n", dir, err)
		return
	}
	entries = dirFilter(entries)

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name() < entries[j].Name()
	})

	for _, e := range entries {
		name := e.Name()
		full := filepath.Join(dir, name)

		info, err := os.Lstat(full)
		if err != nil {
			fmt.Fprintf(os.Stderr, "gols: %s: %v\n", full, err)
			continue
		}

		if useColor {
			if info.IsDir() {
				Blue.ColorPrint(w, name)
				fmt.Fprintln(w)
				continue
			}
			if isExecutableReg(info.Mode()) {
				Green.ColorPrint(w, name)
				fmt.Fprintln(w)
				continue
			}
		}
		fmt.Fprintln(w, name)
	}
}

func isExecutableReg(m os.FileMode) bool {
	return m.IsRegular() && (m&0111) != 0
}
